package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product, listItemIndex := findProductByID(productID)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodGet:
		JSON, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(JSON)

	case http.MethodPut:
		var updated Product
		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updated.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		product = &updated
		productList[listItemIndex] = *product
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
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

	case http.MethodPost:
		var newProduct Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return

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

func main() {
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler)
	http.ListenAndServe(":5000", nil)
}
