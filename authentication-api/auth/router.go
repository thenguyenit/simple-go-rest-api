package auth

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

var routes = Routes{
	//Authentication
	Route{"Authenticate", "POST", "/auth", Handler.Authenticate},
}

//NewRouter will init router for store package
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, r := range routes {
		router.Methods(r.Method).
			Name(r.Name).
			Path(r.Pattern).
			Handler(r.HandlerFunc)
		fmt.Println(r)
	}

	return router
}
