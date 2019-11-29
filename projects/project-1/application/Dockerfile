FROM golang:1.13 AS builder
LABEL maintainer misawa_koyuta@cyberagent.co.jp

WORKDIR /go/src/project-1/application
COPY . .
RUN make build-all
ENTRYPOINT ["make"]

FROM alpine AS base
RUN apk update \
 && apk add --no-cache ca-certificates \
 && apk add --no-cache tzdata \
 && rm -rf /var/cache/apk/*

FROM base AS api
COPY --from=builder /go/src/project-1/application/bin/api /usr/local/bin/api
ENTRYPOINT ["/usr/local/bin/api"]
CMD ["--help"]
