package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Adds the appropriate header to respond with an http file, and prints the response (with header) into the writer
func respondWithHtmlFile(w http.ResponseWriter, r *http.Request, path string) {
	w.Header().Set("Content-Type", "text/html")
	dat, err := ioutil.ReadFile(path)
	if (err != nil) {
		panic(err)
	}
	fmt.Fprintf(w, string(dat))
}
