package auth

import (
	"github.com/thenguyenit/simple-go-rest-api/product-api/db"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	Collection Collection
}

const dbName = "dummyStore"

type Collection struct {
	Database string
	Name     string
}

var repo = Repository{
	Collection: Collection{
		Database: dbName,
		Name:     "user",
	},
}

func (r *Repository) ValidateUser(user User) User {
	var result User
	session := db.NewSession()
	c := session.DB(r.Collection.Database).C(r.Collection.Name)
	c.Find(bson.M{"username": user.Username, "password": user.Password}).One(&result)

	return result
}
