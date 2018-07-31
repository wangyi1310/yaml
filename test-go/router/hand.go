package main

import (
	"encoding/json"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// composite response body
	var res = map[string]string{"result": "succ", "name": "hello"}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
