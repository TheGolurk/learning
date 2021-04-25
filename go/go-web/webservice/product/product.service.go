package product

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

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
		productList := getProductList()
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

		_, err = addOrUpdateProduct(newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return

	}
}
