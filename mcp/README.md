# Introduction

This project is used to quickly create a project based on the MCP SDK and presets some basic configurations.

# Download template locally with gonew

Install gonew if you have not already.

```bash
go install github.com/betterde/gonew@latest
```

Download this template locally:

```bash
gonew github.com/betterde/template/mcp your.domain/module
```

# Config

```yaml
env: production

http:
  listen: 0.0.0.0:8443
logging:
  level: ERROR
```

# Environment

```env
PREFIX=
# Streamable HTTP Configuration
${PREFIX}_HTTP_LISTEN=0.0.0.0:8080
${PREFIX}_LOGGING_LEVEL=DEBUG
```
