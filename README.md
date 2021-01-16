# girie

***girie*** ("go" + "kirie") is a tool for data/metadata extraction from HTML or URL.

### Main goal:

* To have a microservice with API ([GraphQL](https://en.wikipedia.org/wiki/GraphQL)) for ETL pipelines.  
* Provide a plugin endpoint to other tool - [gosquito](https://github.com/livelace/gosquito).

### Features:

* Extract the primary article ([boilerpipe](https://github.com/kohlschutter/boilerpipe), [go-domdistiller](https://github.com/markusmobius/go-domdistiller)) from a web page (HTML and text).
* Extract [JSON-LD](https://en.wikipedia.org/wiki/JSON-LD).
* Extract [Microdata](https://en.wikipedia.org/wiki/Microdata_(HTML)).
* Extract [Opengraph](https://en.wikipedia.org/wiki/Facebook_Platform#Open_Graph_protocol).
* Extract [RDFa](https://en.wikipedia.org/wiki/RDFa).
* Extract images from an entire page or from a page's article.

### Quick start:

```shell script
# start daemon
user@localhost ~ $ docker run --name girie -ti --rm docker.io/livelace/girie:v1.0.0
INFO[16.01.2021 11:38:59.101] girie v1.0.0      
WARN[16.01.2021 11:38:59.102] config error       error="Config File \"config.toml\" Not Found in \"[/etc/girie]\""
INFO[16.01.2021 11:38:59.102] listen :8080 

# execute query
user@localhost ~ $ docker exec girie curl --location --request POST 'http://127.0.0.1:8080/?retry=3&timeout=3' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"{data(url:\"https://iz.ru/1091344/2020-11-24/effektivnost-vaktciny-sputnik-v-prevysila-95\"){html,url,jsonld,microdata,opengraph,rdfa,article{images{alt,src},text},page{images{alt,src},text}}}"}' | jq  
```

### Config example:

```toml
[default]

# Options priority order (top -> down):
# 1. Configuration file.
# 2. Environment variables.
# 3. Query options.

# env GIRIE_LISTEN=":8080"
# listen = ":8080"

# env: GIRIE_PROXY="http://127.0.0.1:3128"
# url: http://127.0.0.1:8080/?proxy="http://127.0.0.1:3128"
# proxy = "http://127.0.0.1:3128"

# env: GIRIE_RETRY=2
# url: http://127.0.0.1:8080/?retry=2
# retry = 2

# env: GIRIE_TIMEOUT=2
# url: http://127.0.0.1:8080/?timeout=2
# timeout = 2

# env: GIRIE_USER_AGENT="girie v1.0.0"
# url: http://127.0.0.1:8080/?user_agent="curl 3000"
# user_agent = "girie v1.0.0"
```
