package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func findProductByID(productID int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}
	return nil, 0
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Println("finished", time.Since(start))
	})
}

func main() {
	productsListHandler := http.HandlerFunc(productsHandler)
	productsItemHandler := http.HandlerFunc(productHandler)

	http.Handle("/products", middlewareHandler(productsListHandler))
	http.Handle("/products/", middlewareHandler(productsItemHandler))
	http.ListenAndServe(":5000", nil)
}