package handlers

import (
	"net/http"
	"strconv"

	"github.com/geoff-maddock/golang-microservice/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//  201: noContent

// DeletePRoducts deletes a product form the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.log.Println("Handle DELETE Product ", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.DeleteProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
