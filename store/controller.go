package store

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct{}

var controller Controller

func (c *Controller) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := repo.GetProducts()
	json, _ := json.Marshal(products)

	w.Write(json)
}

func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
	//read body request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error AddProduct", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestedProduct Product
	json.Unmarshal(body, requestedProduct)

	log.Println(requestedProduct)

}
