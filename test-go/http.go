package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello, ", html.EscapeString(request.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":44441", nil))
}
