// +build dev

package ui

import (
	"net/http"
)

// FS using the native file system restricted to a specific directory tree
var FS http.FileSystem = http.Dir(
	ImportPathToDir("github.com/mozey/vfsgen-example/static/ui/data"))

