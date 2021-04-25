package product

type Product struct {
	ProductID      int    `json:"product_id"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"price_per_unit"`
	QuantityOnHand int    `json:"quantity_on_hand"`
	ProductName    string `json:"product_name"`
}
