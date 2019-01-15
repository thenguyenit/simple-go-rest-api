package store

type Product struct {
	ID     int     `bson:"_id"`
	Title  string  `bson:"title"`
	Image  string  `bson:"image"`
	Price  float64 `bson:"price"`
	Rating int8    `bson:"rating"`
}

type Products []Product
