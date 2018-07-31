package main

import (
	"fmt"
	"net/http"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func main() {
	http.HandleFunc("/", Indexhandler)
	http.ListenAndServe(":8881", nil)
}
