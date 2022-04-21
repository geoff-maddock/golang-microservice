package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	// handle a creation of a resource
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// handle an update to a resource
	if r.Method == http.MethodPut {
		p.log.Println("getting to put")
		// expect the id in the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		matches := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(matches) != 1 {
			p.log.Println("Invalid URI more than one id")
			http.Error(rw, "Invlid URI", http.StatusBadRequest)
			return
		}

		if len(matches[0]) != 2 {
			p.log.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := matches[0][1]
		id, _ := strconv.Atoi(idString)

		p.log.Println("got id", id)

		p.updateProduct(id, rw, r)
	}

	// if no method was supplied
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {

	p.log.Println("Handle post logging")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
	}

	data.AddProduct(prod)
	p.log.Printf("Prod: %#v", prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle PUT update")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
