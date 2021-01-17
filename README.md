# girie

***girie*** ("go" + "kirie") is a tool for data/metadata extraction from web pages.

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
# Start daemon:
user@localhost ~ $ docker run --name girie -ti --rm docker.io/livelace/girie:v1.2.0
INFO[16.01.2021 11:38:59.101] girie v1.2.0      
WARN[16.01.2021 11:38:59.102] config error       error="Config File \"config.toml\" Not Found in \"[/etc/girie]\""
INFO[16.01.2021 11:38:59.102] listen :8080 

# Get API IP:
SERVER=`docker inspect -f "{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}" girie`


# GET + URL:
user@localhost ~ $ docker exec girie curl -s -L -g --request GET \
'http://127.0.0.1:8080/api/?query={data(url:"https://iz.ru/1091344/2020-11-24/effektivnost-vaktciny-sputnik-v-prevysila-95"){article{text_spans_block}}}' | jq  



# POST + URL:
QUERY=`cat << EOF
{
    "query": "{
        data(url: \"https://iz.ru/1091344/2020-11-24/effektivnost-vaktciny-sputnik-v-prevysila-95\") {
            article{
                images{alt,src}
            }
        }
    }"
}
EOF
`

QUERY=`echo $QUERY | tr -d " \n"`

curl -s -L -X POST "http://${SERVER}:8080/api/?retry=3&timeout=3" \
--header "Content-Type: application/json" \
--data-raw "${QUERY}" | jq  



# POST + HTML:
BASE64=`curl -s "https://iz.ru/1091344/2020-11-24/effektivnost-vaktciny-sputnik-v-prevysila-95" | base64 -w0`

QUERY=`cat << EOF
{
    "query": "{
        data(html: \"${BASE64}\") {
            article{
                images{alt,src},
                text,
                text_spans,
                text_spans_append,
                text_spans_block,
            },
            html,
            url,
            jsonld,
            microdata,
            opengraph,
            page{
                images{alt,src},
                text
            }
            rdfa
        }
    }"
}
EOF
`

echo $QUERY | tr -d " \n" > "/tmp/query.json"

curl -s -L -X POST "http://${SERVER}:8080/api/?retry=3&timeout=3" \
--header "Content-Type: application/json" \
--data "@/tmp/query.json" | jq  
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
# url: http://127.0.0.1:8080/api/?proxy="http://127.0.0.1:3128"
# proxy = "http://127.0.0.1:3128"

# env: GIRIE_RETRY=2
# url: http://127.0.0.1:8080/api/?retry=2
# retry = 2

# env: GIRIE_TIMEOUT=2
# url: http://127.0.0.1:8080/api/?timeout=2
# timeout = 10

# env: GIRIE_USER_AGENT="girie v1.2.0"
# url: http://127.0.0.1:8080/api/?user_agent="curl 3000"
# user_agent = "girie v1.2.0"
```
