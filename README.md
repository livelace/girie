# girie

***girie*** ("go" + "kirie") is a tool for data and metadata extraction from raw HTML or URL.

### Main goal:

* To have a microservice with API ([GraphQL](https://en.wikipedia.org/wiki/GraphQL)) for ETL pipelines.  
* Provide a plugin endpoint to other tool - [gosquito](https://github.com/livelace/gosquito).

### Features:

* Extract the primary article ([boilerpipe](https://github.com/kohlschutter/boilerpipe), [go-domdistiller](https://github.com/markusmobius/go-domdistiller)) from a web page (HTML and text).
* Extract [JSON-LD](https://en.wikipedia.org/wiki/JSON-LD).
* Extract [Microdata](https://en.wikipedia.org/wiki/Microdata_(HTML)).
* Extract [Opengraph](https://en.wikipedia.org/wiki/Facebook_Platform#Open_Graph_protocol).
* Extract [RDFa](https://en.wikipedia.org/wiki/RDFa).
* Extract images from an entire page or from an article.

### Quick start:

```shell script
# Docker:
user@localhost /tmp $ docker run -ti --rm livelace/girie:v1.0.0
```

### Example:

```shell
curl --location --request POST 'http://127.0.0.1:8080/?retry=3&timeout=3' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"{data(url:\"https://iz.ru/1091344/2020-11-24/effektivnost-vaktciny-sputnik-v-prevysila-95\"){html,url,jsonld,microdata,opengraph,rdfa,article{images{alt,src},text},page{images{alt,src},text}}}"}' | jq
```

### Config example:

```toml
[default]

# env GIRIE_LISTEN=":8080"
# listen = ":8080"

# env GIRIE_PROXY="http://127.0.0.1:3128"
# proxy = "http://127.0.0.1:3128"

# env GIRIE_RETRY=2
# retry = 2

# env GIRIE_TIMEOUT=2
# timeout = 2

# env GIRIE_USER_AGENT="girie v1.0.0"
# user_agent = "girie v1.0.0"
```
