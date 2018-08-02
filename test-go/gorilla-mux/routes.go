package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	// namespace routes
	Route{
		"getnamee",
		"GET",
		"/api/v1/getname",
		GetName,
	},
	Route{
		"getnum",
		"GET",
		"/api/v1/getnum",
		GetNum,
	},
} //用：的化前提是已经知道类型了如果不知道的情况下需要初始化可以这样搞
