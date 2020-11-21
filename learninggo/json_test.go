package learninggo

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func TestJsonMarshal(t *testing.T) {
	product := &Product{
		ProductID:      123,
		Manufacturer:   "BB Company",
		PricePerUnit:   "12.99",
		Sku:            "43534cere45",
		Upc:            "4144566454",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(productJSON))
}

func TestJsonUnMarshal(t *testing.T) {

	productJSON := `{
		"productId":123,
		"manufacturer":"BB Company",
		"sku":"43534cere45",
		"upc":"4144566454",
		"pricePerUnit":"12.99",
		"quantityOnHand":28,
		"productName":"Gizmo"		
		}`

	product := Product{}
	err := json.Unmarshal([]byte(productJSON), &product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(product)
	fmt.Printf("Product '%s' unmarshalled\n", product.ProductName)
}

func TestJsonEncode(t *testing.T) {
	// ToDo

}
