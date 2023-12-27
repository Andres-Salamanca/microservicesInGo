package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct{

	ID int `json:"id"`
	NAME string `json:"name" validate:"required"`
	DESCRIPTION string `json:"description"`
	PRICE float32 `json:"price" validate:"gt=0"`
	SKU string `json:"sku" validate:"sku"`
	CREATEDOn string `json:"-"`
	UPDATEDOn string `json:"-"`
	DELETEDOn string `json:"-"`

}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {

	re:= regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(),-1)

	if len(matches)!=1{
		return false
	}
	return true
}

func (p *Products) ToJson(w io.Writer)error{

	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader)error{
	e := json.NewDecoder(r)
	return e.Decode(p)

}

func GetProducts() Products{
	return productList
}

func UpdateProduct(id int, p*Product)error{

	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil



}


func AddProduct(p *Product){
	p.ID = getNextId()
	productList = append(productList, p)

}

func getNextId()int{
	lp := productList[len(productList)-1]
	return lp.ID +1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int)(*Product, int, error){
	for i, p := range productList{

		if p.ID == id {
			return p,i,nil
			
		}


	}

	return nil,-1,ErrProductNotFound 
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