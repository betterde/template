# Introduction

This project is used to quickly create a project based on the Fiber framework and presets some basic configurations.

# Download template locally with gonew

Install gonew if you have not already.

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

Download this template locally:

```bash
gonew github.com/betterde/template/fiber your.domain/module
```

# Config

```yaml
env: production

http:
  listen: 0.0.0.0:8443
  tlsKey: /certs/domain.tld.key
  tlsCert: /certs/domain.tld.crt
logging:
  level: ERROR
```

# Environment

```env
PREFIX=
# General configration
${PREFIX}_ENV=production
${PREFIX}_LOGGING_LEVEL=INFO

# API configuration
${PREFIX}_HTTP_LISTEN=0.0.0.0:443

# TLS File provider
${PREFIX}_HTTP_TLSKEY=/certs/domain.tld.key
${PREFIX}_HTTP_TLSCERT=/certs/domain.tld.crt
```