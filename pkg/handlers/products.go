package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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
	} else if r.Method == http.MethodPut {
		// expect the URI containing ID e.g. /3, /andy
		p.l.Println(r.URL.Path)
		uri := strings.Split(r.URL.Path, "/")

		if len(uri) != 3 || uri[0] != "" {
			p.l.Println("Length of URI should be 2")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if uri[1] != "products" {
			p.l.Println("First uri should be /products")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(uri[2])
		if err != nil {
			p.l.Println("The last uri unable to convert to number", uri[2])
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, rw, r)
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
		return
	}

	data.CreateProduct(newProduct)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	// curl  localhost:9090 -X POST -d '{"id": 5, "name": "iPad Air 5", "desc": "A new product", "price": 15000, "sku": "test123"}'
	p.l.Println("Handle PUT products")

	newProduct := &data.Product{}
	if err := newProduct.FromJSON(r.Body); err != nil {
		http.Error(rw, "Loading error from the JSON of product.", http.StatusBadRequest)
		return
	}

	err := data.UpdateProduct(id, newProduct)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found by given ID", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
