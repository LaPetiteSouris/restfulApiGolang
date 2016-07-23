package main

import (
	"encoding/json"
	"net/http"
)

func response(rw http.ResponseWriter, request *http.Request) {
	response := map[string]interface{}{
		"request": "done",
	}
	jsonResponse, _ := json.Marshal(response)
	rw.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/ping", response)
	http.ListenAndServe(":3000", nil)
}
