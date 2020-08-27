// +build vfs

package main

import (
	"log"
	"github.com/shurcooL/vfsgen"
	"net/http"
	"github.com/mozey/vfsgen-example/static/ui"
)

func main() {
	var FS http.FileSystem = http.Dir(
		ui.ImportPathToDir("github.com/mozey/vfsgen-example/static/ui/data"))

	err := vfsgen.Generate(FS, vfsgen.Options{
		PackageName:  "ui",
		BuildTags:    "!dev",
		VariableName: "FS",
		Filename:     "static/ui/vfs.go",
	})

	if err != nil {
		log.Fatalln(err)
	}
}
