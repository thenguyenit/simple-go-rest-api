package product

import "github.com/thenguyenit/simple-go-rest-api/product-api/db"

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
		Name:     "store",
	},
}

func (r *Repository) GetProducts() Products {
	var products Products
	session := db.NewSession()
	c := session.DB(r.Collection.Database).C(r.Collection.Name)
	c.Find(nil).All(&products)

	return products
}

func (r *Repository) AddProduct(p Product) (Product, error) {
	session := db.NewSession()
	c := session.DB(r.Collection.Database).C(r.Collection.Name)
	err := c.Insert(p)
	if err != nil {
		return p, err
	}
	return p, err
}

func (r Repository) UpdateProduct(p Product) bool {
	session := db.NewSession()
	c := session.DB(r.Collection.Database).C(r.Collection.Name)
	err := c.UpdateId(p.ID, p)
	if err != nil {
		return false
	}
	return true
}

func (r Repository) DeleteProduct(pID int) bool {
	session := db.NewSession()
	c := session.DB(r.Collection.Database).C(r.Collection.Name)
	err := c.RemoveId(pID)
	if err != nil {
		return false
	}
	return true
}
