package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := Router()
	fmt.Println("Server is online and serving on port 3000")
	http.ListenAndServe(":3000", router)
}
