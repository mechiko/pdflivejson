package embedded

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed example.json
var JsonExample []byte

//go:embed RobotoCondensed-Regular.ttf
var Regular []byte

//go:embed RobotoCondensed-Bold.ttf
var Bold []byte

//go:embed RobotoCondensed-Italic.ttf
var Italic []byte

//go:embed RobotoCondensed-BoldItalic.ttf
var BoldItalic []byte

//go:embed assets
var EmbeddedAssets embed.FS

//go:embed root
var Root embed.FS

func GetFileSystem() http.FileSystem {
	fsys, err := fs.Sub(Root, "root")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
