// +build dev

package main

import (
	"fmt"
	"github.com/mozey/vfsgen-example/static/ui"
	"io/ioutil"
	"path"

)

func main() {
	f, err := ui.FS.Open(path.Join("docs", "index.html"))
	if err != nil {
		panic(err)
	}
	defer (func() {
		_ = f.Close()
	})()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", buf)
}
