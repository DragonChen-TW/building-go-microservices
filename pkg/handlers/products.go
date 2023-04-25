package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	// curl localhost:9090/products | jp
	p.l.Println("Handle GET products")

	productList := data.GetProducts()
	return c.JSON(http.StatusOK, productList)
}

func (p *Products) CreateProduct(c echo.Context) error {
	/*
		curl localhost:9090/products -X POST \
		-H 'Content-Type: application/json' \
		-d '{"id": 5, "name": "iPad Air 5", "SKU": "apple-ipad-five", "desc": "A new product", "price": 15000}'
	*/
	p.l.Println("Handle POST products")

	newProduct := data.NewProduct()
	if err := c.Bind(newProduct); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := newProduct.Validate(); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	p.l.Println("new", newProduct)

	data.CreateProduct(newProduct)
	return nil
}

func (p *Products) UpdateProduct(c echo.Context) error {
	/*
		curl localhost:9090/products/3 -X PUT \
		-H 'Content-Type: application/json' \
		-d '{"name": "iPad Air 5", "SKU": "ipad-air-five", "desc": "Amazing iPad!!!!!!", "price": 12000}'
	*/
	p.l.Println("Handle PUT products")

	newProduct := &data.Product{}
	if err := c.Bind(newProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := newProduct.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	p.l.Println("update", newProduct)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	err = data.UpdateProduct(id, newProduct)
	if err == data.ErrProductNotFound {
		c.JSON(http.StatusNotFound, "Product not found by given ID")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Product not found")
	}

	return nil
}

// // Process is the middleware function.
// func ValidateNewProductMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		newProduct := data.NewProduct()
// 		if err := c.Bind(newProduct); err != nil {
// 			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 		}
// 		if err := newProduct.Validate(); err != nil {
// 			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 		}
// 		fmt.Println(newProduct)
// 		return next(c)
// 	}
// }
