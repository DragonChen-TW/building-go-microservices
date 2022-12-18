package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dragonchen-tw/building-go-microservices/pkg/handlers"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Building Go microservices
// Date:	2022/12/18

func main() {
	l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	// our HelloHandler fulfill the interface of Handler
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe("127.0.0.1:9090", sm)
}
