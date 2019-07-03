package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		http.MethodGet,
		"/accounts/{accountId}",
		GetAccount,
	},
	Route{
		"HealthCheck",
		http.MethodGet,
		"/health",
		HealthCheck,
	},
	Route{
		"Testability",
		http.MethodPut,
		"/testability/healthy/{state}",
		SetHealthyState,
	},
}
