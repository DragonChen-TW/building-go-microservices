package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/dragonchen-tw/building-go-microservices/pkg/handlers"
	"github.com/labstack/echo/v4"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Building Go microservices
// Date:	2022/12/18

func main() {
	// l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)

	e := echo.New()
	e.Server.ReadTimeout = 1 * time.Second
	e.Server.WriteTimeout = 1 * time.Second
	e.Server.IdleTimeout = 120 * time.Second

	ph := handlers.NewProducts(e.StdLogger)

	g := e.Group("/products")
	g.GET("", ph.GetProducts).Name = "get-products"
	g.POST(
		"", ph.CreateProduct,
		// handlers.ValidateNewProductMiddleware,
	).Name = "create-product"
	g.PUT(
		"/:id", ph.UpdateProduct,
		// handlers.ValidateNewProductMiddleware,
	).Name = "update-product"

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start("127.0.0.1:9090"))
}
