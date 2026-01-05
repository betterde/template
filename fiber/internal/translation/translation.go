package translation

import "embed"

//go:embed locales/*.json
var locales embed.FS
