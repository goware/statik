//go:generate statik -f -src=./public

package main

import (
	"log"
	"net/http"

	"github.com/goware/statik/example/static"
	"github.com/goware/statik/fs"
)

// Before buildling, run go generate.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func main() {
	statikFS, err := fs.New(static.Asset)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}
