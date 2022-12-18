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
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("test Products.")

	productList := data.GetProducts()
	if err := productList.ToJSON(rw); err != nil {
		http.Error(rw, "Dumping error for the JSON of productList.", http.StatusInternalServerError)
	}
}
