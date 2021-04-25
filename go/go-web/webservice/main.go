package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductID      int    `json:"product_id"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"price_per_unit"`
	QuantityOnHand int    `json:"quantity_on_hand"`
	ProductName    string `json:"product_name"`
}

var productList []Product

func init() {
	productJSON := `[
		{
			"product_id": 1,
			"manufacturer": "some",
			"sku": "p5fsd12123",
			"upc": "1231312313",
			"price_per_unit": "12.312",
			"quantity_on_hand": 123,
			"product_name": "nameXD"
		}
	]`
	err := json.Unmarshal([]byte(productJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(productsJson)
	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
