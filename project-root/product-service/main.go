package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var products = []Product{
	{ID: "1", Name: "Laptop"},
	{ID: "2", Name: "Smartphone"},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/products", getProducts)
	http.ListenAndServe(":8082", nil)
}
