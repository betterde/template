# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine AS builder

WORKDIR /src

ARG TARGETOS
ARG TARGETARCH
ARG APP_NAME={{YOUR_SERVER}}-mcp-server
ARG APP_DESC="A Go-based MCP server for {{YOUR_SERVER}}."
ARG VERSION=develop
ARG BUILD=current
ARG COMMIT=none

RUN apk add --no-cache ca-certificates tzdata

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-$(go env GOARCH)} go build \
    -trimpath \
    -ldflags="-s -w \
    -X 'github.com/{{YOUR_NAMESPACE}}/internal/build.Version=${VERSION}' \
    -X 'github.com/{{YOUR_NAMESPACE}}/internal/build.Build=${BUILD}' \
    -X 'github.com/{{YOUR_NAMESPACE}}/internal/build.Commit=${COMMIT}'" \
    -o /out/{{YOUR_SERVER}}-mcp-server .

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S mysql-mcp-server \
    && adduser -S -G mysql-mcp-server mysql-mcp-server

WORKDIR /etc/mcp-servers/{{YOUR_SERVER}}

COPY --from=builder /out/{{YOUR_SERVER}}-mcp-server /usr/local/bin/{{YOUR_SERVER}}-mcp-server

USER {{YOUR_SERVER}}-mcp-server:{{YOUR_SERVER}}-mcp-server

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/{{YOUR_SERVER}}-mcp-server"]
CMD ["serve"]
