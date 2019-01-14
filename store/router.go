package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Imports

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddProduct",
		"POST",
		"/AddProduct",
		controller.Index,
	},
}

// More routes.....

// NewRouter function configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

// // Get Authentication token GET /
// func (c *Controller) GetToken(w http.ResponseWriter, req *http.Request) {
// 	var user User
// 	_ = json.NewDecoder(req.Body).Decode(&user)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": user.Username,
// 		"password": user.Password,
// 	})

// 	log.Println("Username: " + user.Username)
// 	log.Println("Password: " + user.Password)

// 	tokenString, error := token.SignedString([]byte("secret"))
// 	if error != nil {
// 		fmt.Println(error)
// 	}
// 	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
// }

// /* Middleware handler to handle all requests for authentication */
// func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		authorizationHeader := req.Header.Get("authorization")
// 		if authorizationHeader != "" {
// 			bearerToken := strings.Split(authorizationHeader, " ")
// 			if len(bearerToken) == 2 {
// 				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
// 					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 						return nil, fmt.Errorf("There was an error")
// 					}
// 					return []byte("secret"), nil
// 				})
// 				if error != nil {
// 					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
// 					return
// 				}
// 				if token.Valid {
// 					log.Println("TOKEN WAS VALID")
// 					context.Set(req, "decoded", token.Claims)
// 					next(w, req)
// 				} else {
// 					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
// 				}
// 			}
// 		} else {
// 			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
// 		}
// 	})
// }
