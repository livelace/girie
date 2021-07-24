FROM            harbor-core.k8s-2.livelace.ru/infra/service-core:latest

COPY            "girie" "/usr/local/bin/girie"

# user and group.
RUN             groupadd -g 1000 "user" && \
                useradd -l -u 1000 -g "user" -s "/bin/bash" -m "user"

USER            "user"

CMD             ["/usr/local/bin/girie"]
