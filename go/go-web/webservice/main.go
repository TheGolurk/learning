package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thegolurk/inventoryservice/database"
	"github.com/thegolurk/inventoryservice/product"
)

const apiBasePath = "/api"

func main() {
	database.SetUpDatabase()
	product.SetUpRoutes(apiBasePath)

	http.ListenAndServe(":5000", nil)
}
