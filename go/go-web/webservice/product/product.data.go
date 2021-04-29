package product

import (
	"database/sql"
	"log"

	"github.com/thegolurk/inventoryservice/database"
)

func getProduct(productID int) (*Product, error) {
	row := database.DbConn.QueryRow(`SELECT * FROM nventorydb.products
	WHERE productId = ?`, productID)

	var product Product
	err := row.Scan(&product.ProductID,
		&product.Manufacturer,
		&product.Sku,
		&product.Upc,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.ProductName)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}

func removeProduct(productID int) error {
	_, err := database.DbConn.Query(`DELETE FROM nventorydb.products WHERE productId = ?`, productID)
	if err != nil {
		return err
	}

	return nil
}

func getProductList() ([]Product, error) {
	results, err := database.DbConn.Query(`SELECT * FROM nventorydb.products`)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	products := make([]Product, 0)
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)
	}

	return products, nil
}

func updateProduct(product Product) error {
	_, err := database.DbConn.Exec(`UPDATE nventorydb.products SET
		manufacturer=?,
		sku=?,
		upc=?,
		pricePerUnit=CAST(? AS DECIMAL(13,2)),
		quantityOnHand=?,
		productName=?
		WHERE productId=?
	`, product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID)

	if err != nil {
		return err
	}
	return nil
}

func insertProduct(product Product) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO nventorydb.products 
						(manufacturer, sku, upc, pricePerUnit, quantityOnHand, productName) VALUES 
						(?, ?, ?, ?, ?, ?) `, product.Manufacturer, product.Sku, product.Upc,
		product.PricePerUnit, product.QuantityOnHand, product.ProductName)

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Print("error result", err)
		return 0, err
	}

	return int(insertID), nil
}
