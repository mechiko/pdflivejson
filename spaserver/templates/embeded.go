package templates

import (
	"embed"
)

//go:embed header footer home index live
var root embed.FS
