package db

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func NewSession() *mgo.Session {
	mgoURL := os.Getenv("MGO_URL")
	session, err := mgo.Dial(mgoURL)
	if err != nil {
		log.Fatal("Connect to " + mgoURL + " is failed")

	}
	return session
}
