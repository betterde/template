ARG VERSION=latest

FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

ARG MODULE=github.com/betterde/server
ARG BINARY_NAME=server
ARG INSTALL_PATH=/usr/local/bin

RUN apk add --update gcc git make

ENV GOPATH=/tmp/buildcache
COPY . /go/src/cdns
WORKDIR /go/src/cdns
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X '${MODULE}/cmd.version=${VERSION}' -X '${MODULE}/cmd.build=`date -u`' -X '${MODULE}/cmd.commit=`git rev-parse HEAD`'" -o bin/${BINARY_NAME} main.go

FROM --platform=$BUILDPLATFORM alpine:latest

LABEL org.opencontainers.image.url="https://github.com/betterde/server"
LABEL org.opencontainers.image.titile="server"
LABEL org.opencontainers.image.vendor="Betterde Inc."
LABEL org.opencontainers.image.source="https://github.com/betterde/server"
LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.authors="George <george@betterde.com>"
LABEL org.opencontainers.image.created="2024-08-21 20:35:00"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.description="An open-source lightweight DNS server that switches to ACME DNS challenge."
LABEL org.opencontainers.image.documentation="https://github.com/betterde/server"

COPY --from=builder /go/src/cdns/bin/cdns /usr/local/bin/cdns
RUN mkdir -p /etc/server
WORKDIR /root/

VOLUME ["/etc/server"]
ENTRYPOINT ["/usr/local/bin/server"]
EXPOSE 443