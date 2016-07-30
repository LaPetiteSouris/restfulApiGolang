package main

import (
	"encoding/json"
	"net/http"
)

func getDataHandler(rw http.ResponseWriter, request *http.Request) {
	response := map[string]interface{}{
		"request":     "done",
		"requestType": request.Method,
	}
	jsonResponse, _ := json.Marshal(response)
	rw.Write(jsonResponse)
}
