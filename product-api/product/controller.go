package product

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct{}

var Handler Controller

func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := repo.GetProducts()
	json, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(json)
}

func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
	//read body request
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var product Product
	if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		w.Write([]byte(err.Error()))
		return
	}

	p, err := repo.AddProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json, _ := json.Marshal(p)
	w.Write(json)
	return
}

func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//read body request
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var product Product
	if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		w.Write([]byte(err.Error()))
		return
	}

	success := repo.UpdateProduct(product)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(product)
	w.Write(json)
	return
}

func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["id"])

	success := repo.DeleteProduct(pID)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
