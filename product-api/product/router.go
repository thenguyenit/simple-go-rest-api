package product

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type MiddlewareFunc func(http.Handler) http.Handler

const pathPrefix = "/api"

var routes = Routes{
	//Product
	Route{"IndexProduct", "GET", "/product", Handler.GetProducts},
	Route{"AddProduct", "POST", "/product/add", Handler.AddProduct},
	Route{"UpdateProduct", "PUT", "/product/update", Handler.UpdateProduct},
	Route{"DeleteProduct", "DELETE", "/product/delete/{id:[0-9]}", Handler.DeleteProduct},
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

	router.Use(ValidateTokenMiddleware)

	return router
}

func ValidateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		tokenString := strings.Replace(bearer, "Bearer ", "", 1)

		if tokenString == "" {
			http.Error(w, "Token is nil", http.StatusForbidden)
			return
		}
		fmt.Println(tokenString)

		hmacSecret := os.Getenv("APP_SECRET")
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSecret, nil
		})

		if token.Valid {
			// We found the token in our map
			log.Printf("Token is valid")
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
			// http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
