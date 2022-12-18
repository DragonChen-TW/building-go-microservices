package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Building Go microservices
// Date:	2022/12/18

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello, World!")

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Got data-reading error.", http.StatusBadRequest)
			return
		}
		// Check the input data
		// log.Printf("Data: %s\n", data)

		fmt.Fprintf(rw, "Hello %s.\n", data)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye, World!")
	})

	http.ListenAndServe("127.0.0.1:9090", nil)
}
