package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

type Controller struct{}

var Handler Controller

func (c *Controller) Authenticate(w http.ResponseWriter, r *http.Request) {
	hmacSecret := os.Getenv("APP_SECRET")
	// buf := new(bytes.Buffer)
	// io.Copy(buf, r.Body)

	var credential User
	err := json.NewDecoder(r.Body).Decode(&credential)

	if err != nil {
		log.Println("Invalid input data")
	}
	r.Body.Close()

	//Validate username/password
	user := repo.ValidateUser(credential)
	if user.ID > 0 {
		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			// "roles":    user.Roles,
			// "nbf": time.Now().Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, _ := token.SignedString([]byte(hmacSecret))
		w.Write([]byte(tokenString))
	} else {
		w.Write([]byte("Invalid credential, Please try again"))
	}

	return
}
