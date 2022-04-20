package handlers

import (
	"log"
	"net/http"

	"github.com/geoff-maddock/golang-microservice/data"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle an update
	if r.Method == http.MethodPut {

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, rh *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
