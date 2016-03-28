package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodosIndex", "GET", "/todos", TodoIndex},
	Route{"TodoShow", "GET", "/todos/{id}", TodoShow},
	Route{"TodoCreate", "POST", "/todos", TodoCreate},
}