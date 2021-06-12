package girie

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/livelace/girie/pkg/girie/core"
	log "github.com/livelace/logrus"
	distiller "github.com/markusmobius/go-domdistiller"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: false,
		ForceColors:            true,
		ForceQuote:             true,
		FullTimestamp:          true,
		TimestampFormat:        core.DEFAULT_LOG_TIME_FORMAT,
		QuoteEmptyFields:       true,
	})
}

// Pass configuration options to query.
func configMiddleware(v *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("proxy", v.GetString(core.VIPER_DEFAULT_PROXY))
		c.Set("retry", v.GetInt(core.VIPER_DEFAULT_RETRY))
		c.Set("timeout", v.GetInt(core.VIPER_DEFAULT_TIMEOUT))
		c.Set("user_agent", v.GetString(core.VIPER_DEFAULT_USER_AGENT))
		c.Next()
	}
}

// Query execution.
func executeQuery(c *gin.Context) {
	var p core.InputData

	// Derive query body from GET/POST.
	if c.Request.Method == "GET" {
		var b bool
		p.Query, b = c.GetQuery("query")
		if !b {
			e := core.Error{
				Code:        http.StatusBadRequest,
				Description: "Request must contain query parameter.",
				Error:       "not enough data",
			}
			c.JSON(http.StatusBadRequest, e)
			return
		}

	} else if c.Request.Method == "POST" {
		err := c.BindJSON(&p)
		if err != nil {
			e := core.Error{
				Code:        http.StatusBadRequest,
				Description: "Cannot parse JSON data.",
				Error:       err.Error(),
			}
			c.JSON(http.StatusBadRequest, e)
			return
		}

	} else {
		e := core.Error{
			Code:        http.StatusBadRequest,
			Description: "Only GET and POST are allowed.",
			Error:       "method not allowed",
		}
		c.JSON(http.StatusBadRequest, e)
		return
	}

	// Get provided source (HTML or URL) parameters.
	html, url, err := core.GetArguments(p.Query)
	if err != nil {
		e := core.Error{
			Code:        http.StatusInternalServerError,
			Description: "Cannot extract any query arguments (html, url)",
			Error:       err.Error(),
		}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	htmlOrigin := html

	// Exit if no source.
	if len(html) == 0 && len(url) == 0 {
		e := core.Error{
			Code:        http.StatusBadRequest,
			Description: "Request must not be empty. Query arguments must be provided (html, url).",
			Error:       "not enough data",
		}
		c.JSON(http.StatusBadRequest, e)
		return
	}

	// "url" has priority over "html".
	if len(url) > 0 {
		// Get options for page fetching.
		var proxy string
		configProxy := c.GetString("proxy")
		queryProxy := c.Query("proxy")
		if len(queryProxy) > 0 {
			proxy = queryProxy
		} else {
			proxy = configProxy
		}

		retry := core.GetInt(c.Query("retry"), c.GetInt("retry"))

		timeout := core.GetInt(c.Query("timeout"), c.GetInt("timeout"))

		var userAgent string
		configUserAgent := c.GetString("user_agent")
		queryUserAgent := c.Query("user_agent")
		if len(queryUserAgent) > 0 {
			userAgent = queryUserAgent
		} else {
			userAgent = configUserAgent
		}

		html, err = core.FetchPage(url, proxy, retry, timeout, userAgent)
		if err != nil {
			e := core.Error{
				Code:        http.StatusBadRequest,
				Description: "Cannot fetch page.",
				Error:       err.Error(),
			}
			c.JSON(http.StatusBadRequest, e)
			return
		}
	} else {
		b, err := base64.StdEncoding.DecodeString(html)
		html = string(b)

		if err != nil {
			e := core.Error{
				Code:        http.StatusBadRequest,
				Description: "Cannot decode HTML body.",
				Error:       err.Error(),
			}
			c.JSON(http.StatusBadRequest, e)
			return
		}
	}

	// Extract main article from HTML.
	distillHTML, errHTML := distiller.ApplyForReader(strings.NewReader(html), nil)
	distillText, errText := distiller.ApplyForReader(strings.NewReader(html),
		&distiller.Options{ExtractTextOnly: true, SkipPagination: true})

	if errHTML != nil {
		err = errHTML
	} else if errText != nil {
		err = errText
	} else {
		err = nil
	}

	if err != nil {
		e := core.Error{
			Code:        http.StatusInternalServerError,
			Description: "Cannot extract main article",
			Error:       err.Error(),
		}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if distillHTML == nil || distillText == nil {
		e := core.Error{
			Code:        http.StatusInternalServerError,
			Description: "Cannot extract main article",
			Error:       "null data",
		}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	// Fill context data for invocations from other functions .
	data := core.Data{
		HTML: htmlOrigin,
		URL:  url,

		Article: core.Article{
			HTML: distillHTML.HTML,
			Text: distillText.HTML,
		},
		Page: core.Page{
			HTML: html,
		},
	}

	// Execute Graphql stuff after all preparations.
	result := graphql.Do(graphql.Params{
		Schema:         core.Schema,
		Context:        context.WithValue(context.Background(), "data", data),
		OperationName:  p.Operation,
		RequestString:  p.Query,
		VariableValues: p.Variables,
	})

	if result.HasErrors() {
		e := core.Error{
			Code:        http.StatusBadRequest,
			Description: "Cannot execute query.",
			Error:       result.Errors,
		}
		c.JSON(http.StatusBadRequest, e)
	} else {
		c.JSON(http.StatusOK, result.Data)
	}
}

func RunApp() {
	// Greetings.
	log.Info(fmt.Sprintf("%s %s", core.APP_NAME, core.APP_VERSION))

	// Get config.
	config := core.GetConfig()

	// We are ready to go :)
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("INFO[%s] %s %s %s %d %d %s \"%s\" %s\n",
			param.TimeStamp.Format(core.DEFAULT_LOG_TIME_FORMAT),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.BodySize,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.Use(configMiddleware(config))
	router.Any("/api", executeQuery)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Code":   200,
			"Status": "healthy",
		})
	})

	log.Infof("listen %s", config.GetString(core.VIPER_DEFAULT_LISTEN))

	err := router.Run(config.GetString(core.VIPER_DEFAULT_LISTEN))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("start error")
		os.Exit(1)
	}
}
