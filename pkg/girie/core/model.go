package core

type Article struct {
	HTML   string   `json:"html"`
	Images []string `json:"images"`
	Text   string   `json:"text"`
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

type Image struct {
	Alt string `json:"alt"`
	Src string `json:"src"`
}

type Page struct {
	HTML   string   `json:"html"`
	Images []string `json:"images"`
	Text   string   `json:"text"`
}

type PostData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}
