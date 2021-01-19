package core

type Article struct {
	HTML      string   `json:"html"`
	Images    []string `json:"images"`
	Text      string   `json:"text"`
	TextSpans []string `json:"text_spans"`
}

type Data struct {
	HTML string `json:"html"`
	URL  string `json:"url"`

	JsonLd    string `json:"jsonld"`
	Microdata string `json:"microdata"`
	Opengraph string `json:"opengraph"`
	Rdfa      string `json:"rdfa"`
	Title     string `json:"title"`

	Article Article `json:"article"`
	Page    Page    `json:"page"`
}

type Error struct {
	Code        int
	Description string
	Error       interface{}
}

type InputData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

type Image struct {
	Alt    string `json:"alt"`
	Height int    `json:"height"`
	Src    string `json:"src"`
	Width  int    `json:"width"`
}

type Page struct {
	HTML   string   `json:"html"`
	Images []string `json:"images"`
	Text   string   `json:"text"`
}

type TextSpan struct {
	Lang         string `json:"lang"`
	Text         string `json:"text"`
	TokensAmount int    `json:"tokens_amount"`
}
