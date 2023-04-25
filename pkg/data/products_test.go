package data

import (
	"testing"
)

func TestEmptyProduct(t *testing.T) {
	p := NewProduct()
	err := p.Validate()
	if err == nil {
		t.Errorf("Empty product can not pass the validator, %s", err)
	}
}

func TestNoSKUProduct(t *testing.T) {
	p := &Product{
		Name:  "isu",
		Price: 10.00,
	}
	err := p.Validate()
	if err == nil {
		t.Errorf("Product must have SKU, %s", err)
	}
}

func TestSKUFormatProduct(t *testing.T) {
	p := &Product{
		Name:  "isu",
		Price: 10.00,
		SKU:   "abc-sdf-1234",
	}
	err := p.Validate()
	if err == nil {
		t.Errorf("SKU format is wrong, %s", err)
	}
}

func TestCorrectProduct(t *testing.T) {
	p := &Product{
		Name:  "isu",
		Price: 10.00,
		SKU:   "abc-sdf-dddd",
	}
	err := p.Validate()
	if err != nil {
		t.Errorf("SKU format is correct, %s", err)
	}
}
