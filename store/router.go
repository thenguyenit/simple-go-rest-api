package store

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const pathPrefix = "/product"

var routes = Routes{
	Route{
		"IndexProduct",
		"GET",
		"/",
		controller.ProductsHandler,
	},
	Route{
		"AddProduct",
		"POST",
		"/add",
		controller.AddProduct,
	},
	Route{
		"SearchProduct",
		"GET",
		"/search",
		controller.ProductsHandler,
	},
}

//NewRouter will init router for store package
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, r := range routes {
		router.Methods(r.Method).
			Name(r.Name).
			PathPrefix(pathPrefix).
			Path(r.Pattern).
			Handler(r.HandlerFunc)
		fmt.Println(r)
	}

	return router
}
