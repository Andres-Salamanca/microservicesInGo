package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	Data "github.com/Andres-Salamanca/microcourse/Data"
	data "github.com/Andres-Salamanca/microcourse/Data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}



func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {

	d := Data.GetProducts()
	err := d.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request){

	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	
	data.AddProduct(&prod)
	
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle put product",id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = Data.UpdateProduct(id,&prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}



}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(handler http.Handler) http.Handler{

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}

		err := prod.FromJson(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		err=prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		handler.ServeHTTP(rw, r)

	})

		
	
}
