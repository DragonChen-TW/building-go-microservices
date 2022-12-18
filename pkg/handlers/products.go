package handlers

import (
	"log"
	"net/http"

	"github.com/dragonchen-tw/building-go-microservices/pkg/data"
	"github.com/labstack/echo/v4"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(c echo.Context) error {
	// curl  localhost:9090 | jp
	p.l.Println("Handle GET products")

	productList := data.GetProducts()
	return c.JSON(http.StatusOK, productList)
}

// func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {
// 	// curl  localhost:9090 -X POST -d '{"id": 5, "name": "iPad Air 5", "desc": "A new product", "price": 15000, "sku": "test123"}'
// 	p.l.Println("Handle POST products")

// 	newProduct := &data.Product{}
// 	if err := newProduct.FromJSON(r.Body); err != nil {
// 		http.Error(rw, "Loading error from the JSON of product.", http.StatusBadRequest)
// 		return
// 	}

// 	data.CreateProduct(newProduct)
// }

// func (p *Products) UpdateProduct(id int, rw http.ResponseWriter, r *http.Request) {
// 	// curl  localhost:9090 -X POST -d '{"id": 5, "name": "iPad Air 5", "desc": "A new product", "price": 15000, "sku": "test123"}'
// 	p.l.Println("Handle PUT products")

// 	newProduct := &data.Product{}
// 	if err := newProduct.FromJSON(r.Body); err != nil {
// 		http.Error(rw, "Loading error from the JSON of product.", http.StatusBadRequest)
// 		return
// 	}

// 	err := data.UpdateProduct(id, newProduct)
// 	if err == data.ErrProductNotFound {
// 		http.Error(rw, "Product not found by given ID", http.StatusNotFound)
// 		return
// 	} else if err != nil {
// 		http.Error(rw, "Product not found", http.StatusInternalServerError)
// 		return
// 	}
// }
