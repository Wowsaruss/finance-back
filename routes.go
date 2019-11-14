package main

import "net/http"

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"homePage",
		"GET",
		"/api",
		homePage,
	},
	Route{
		"getType",
		"GET",
		"/api/type/{type}",
		getType,
	},
	Route{
		"insertData",
		"POST",
		"/api/insert",
		insertData,
	},
}
