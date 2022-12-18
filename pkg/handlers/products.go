package handlers

import (
	"log"
	"net/http"

	"github.com/dragonchen-tw/building-go-microservices/pkg/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	} else if r.Method == http.MethodPost {
		p.createProduct(rw, r)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	// curl  localhost:9090 | jp
	p.l.Println("Handle GET products")

	productList := data.GetProducts()
	if err := productList.ToJSON(rw); err != nil {
		http.Error(rw, "Dumping error for the JSON of productList.", http.StatusInternalServerError)
	}
}

func (p *Products) createProduct(rw http.ResponseWriter, r *http.Request) {
	// curl  localhost:9090 -X POST -d '{"id": 5, "name": "iPad Air 5", "desc": "A new product", "price": 15000, "sku": "test123"}'
	p.l.Println("Handle POST products")

	newProduct := &data.Product{}
	if err := newProduct.FromJSON(r.Body); err != nil {
		http.Error(rw, "Loading error from the JSON of product.", http.StatusBadRequest)
	}

	p.l.Printf("NewProduct: %v", newProduct)

	data.CreateProduct(newProduct)
}
