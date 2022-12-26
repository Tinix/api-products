package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Product represents a product in a store.
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Products is a slice of Product.
type Products []Product

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := Products{
		Product{ID: "1", Name: "Product 1", Price: 9.99},
		Product{ID: "2", Name: "Product 2", Price: 19.99},
		Product{ID: "3", Name: "Product 3", Price: 29.99},
	}

	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product Product
	if id == "1" {
		product = Product{ID: "1", Name: "Product 1", Price: 9.99}
	} else if id == "2" {
		product = Product{ID: "2", Name: "Product 2", Price: 19.99}
	} else if id == "3" {
		product = Product{ID: "3", Name: "Product 3", Price: 29.99}
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	product.ID = "4"

	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	if id != product.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.Name = "Updated Product"
	product.Price = 99.99

	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id != "1" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
