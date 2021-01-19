package core

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/abadojack/whatlanggo"
	"github.com/enbis/rdfa"
	"github.com/go-resty/resty/v2"
	sourceAST "github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/visitor"
	"github.com/iand/microdata"
	"github.com/microcosm-cc/bluemonday"
	"github.com/otiai10/opengraph/v2"
	u "net/url"
	"strconv"
	"strings"
	"time"
)

func GetArguments(query string) (string, string, error) {
	var fields []*sourceAST.Field
	var html string
	var url string

	ast, err := parser.Parse(parser.ParseParams{Source: query})
	if err != nil {
		return html, url, err
	}

	v := &visitor.VisitorOptions{
		Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
			if node, ok := p.Node.(*sourceAST.Field); ok {
				if node.SelectionSet == nil {
					return visitor.ActionNoChange, nil
				}

				fields = append(fields, node)
			}
			return visitor.ActionNoChange, nil
		},
	}

	visitor.Visit(ast, v, nil)

	for _, f := range fields {
		if f.Name.Value == "data" {
			for _, arg := range f.Arguments {
				switch arg.Name.Value {
				case "html":
					html = arg.Value.GetValue().(string)
				case "url":
					url = arg.Value.GetValue().(string)
				}
			}
		}
	}

	return html, url, nil
}

func ExtractImages(html string) []Image {
	images := make([]Image, 0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err == nil {
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			var alt string
			var height string
			var src string
			var width string

			alt, _ = s.Attr("alt")
			height, _ = s.Attr("height")
			src, _ = s.Attr("src")
			width, _ = s.Attr("width")

			images = append(images, Image{
				Alt:    alt,
				Height: GetInt(height, 0),
				Src:    src,
				Width:  GetInt(width, 0),
			})
		})
	}

	return images
}

func ExtractJSONLd(html string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	return doc.Find("script[type=\"application/ld+json\"]").Text()
}

func ExtractMicrodata(html string, url string) string {
	baseUrl, _ := u.Parse(url)
	p := microdata.NewParser(strings.NewReader(html), baseUrl)
	data, _ := p.Parse()
	b, _ := data.JSON()
	return string(b)
}

func ExtractOpengraph(html string) string {
	ogp := &opengraph.OpenGraph{}
	_ = ogp.Parse(strings.NewReader(html))
	b, _ := json.Marshal(ogp)
	return string(b)
}

func ExtractRDFA(html string) string {
	b, _ := rdfa.Extract(html)
	return string(b)
}

func ExtractTitle(html string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	return doc.Find("title").Text()
}

func FetchPage(url, proxy string, retry, timeout int, userAgent string) (string, error) {
	client := resty.New()

	if len(proxy) > 0 {
		client.SetProxy(proxy)
	}
	client.SetRetryCount(retry)
	client.SetTimeout(time.Duration(timeout) * time.Second)
	client.SetHeaders(map[string]string{
		"User-Agent": userAgent,
	})

	resp, err := client.R().Get(url)

	if resp != nil {
		return resp.String(), err
	} else {
		return "", err
	}
}

func GetInt(v string, d int) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return d
	}
	return i
}

func GetTextSpan(s *string) *TextSpan {
	span := TextSpan{
		Text:         strings.TrimSpace(*s),
		TokensAmount: len(strings.Split(*s, " ")),
	}

	detect := whatlanggo.Detect(*s)

	if detect.IsReliable() {
		span.Lang = detect.Lang.Iso6391()
	}

	return &span
}

func SanitizeHTMLTags(html string) string {
	return bluemonday.StrictPolicy().Sanitize(html)
}
