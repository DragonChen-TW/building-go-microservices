package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"desc"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

type ProductList []*Product

func (pl *ProductList) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(productList); err != nil {
		return err
	}
	return nil
}

func GetProducts() ProductList {
	return productList
}

func CreateProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	pl := productList[len(productList)-1]
	return pl.ID + 1
}

// productList is a hard coded list of products for this
// example data source
var productList = ProductList{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
