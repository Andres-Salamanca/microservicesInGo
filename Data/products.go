package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct{

	ID int `json:"id"`
	NAME string `json:"name"`
	DESCRIPTION string `json:"description"`
	PRICE float32 `json:"price"`
	SKU string `json:"sku"`
	CREATEDOn string `json:"-"`
	UPDATEDOn string `json:"-"`
	DELETEDOn string `json:"-"`

}

type Products []*Product

func (p *Products) ToJson(w io.Writer)error{

	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products{
	return productList
}


var productList = []*Product{
	&Product{
		ID:          1,
		NAME:        "Latte",
		DESCRIPTION: "Frothy milky coffee",
		PRICE:       2.45,
		SKU:         "abc323",
		CREATEDOn:   time.Now().UTC().String(),
		UPDATEDOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		NAME:        "Espresso",
		DESCRIPTION: "Short and strong coffee without milk",
		PRICE:       1.99,
		SKU:         "fjd34",
		CREATEDOn:   time.Now().UTC().String(),
		UPDATEDOn:   time.Now().UTC().String(),
	},
}