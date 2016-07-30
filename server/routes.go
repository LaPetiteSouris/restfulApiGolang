package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			HandleFunc(route.Pattern, route.handler).
			Methods(route.Method)
	}
	return router

}

type Route struct {
	Name    string
	Method  string
	Pattern string
	handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"getData",
		"GET",
		"/getdata",
		getDataHandler,
	},
	Route{
		"createData",
		"POST",
		"/createdata",
		getDataHandler,
	},
}
