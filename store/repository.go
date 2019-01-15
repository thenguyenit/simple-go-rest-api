package store

import (
	mgo "gopkg.in/mgo.v2"
)

type Repository struct {
	Collection Collection
}

const DB_NAME = "dummyStore"
const CollectionName = "dummyStore"

const URL = "mongodb://127.0.0.1:27017/"

type Collection struct {
	Database string
	Name     string
}

var repo = Repository{
	Collection: Collection{
		Database: DB_NAME,
		Name:     "store",
	},
}

func (r *Repository) GetProducts() Products {
	var products Products
	session, err := mgo.Dial(URL)
	if err == nil {
		c := session.DB(r.Collection.Database).C(r.Collection.Name)
		c.Find(nil).All(&products)
	}

	return products
}
