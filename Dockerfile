FROM            docker.io/livelace/gentoo:latest

ARG             VERSION

ENV             GIRIE_BIN="/usr/local/bin/girie"
ENV             GIRIE_TEMP="/tmp/girie"
ENV             GIRIE_URL="https://github.com/livelace/girie"

# portage packages.
RUN             emerge -G -q \
                dev-lang/go && \
                rm -rf "/usr/portage/packages"

# build application.
RUN             git clone --depth 1 --branch "$VERSION" "$GIRIE_URL" "$GIRIE_TEMP" && \
                cd "$GIRIE_TEMP" && \
                go build "github.com/livelace/girie/cmd/girie" && \
                cp "girie" "$GIRIE_BIN" && \
                rm -rf "/root/go" "$GIRIE_TEMP"

RUN             useradd -m -u 1000 -s "/bin/bash" "girie"

USER            "girie"

WORKDIR         "/home/girie"

CMD             ["/usr/local/bin/girie"]
