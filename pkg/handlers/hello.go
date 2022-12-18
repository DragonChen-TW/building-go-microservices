package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello, World!")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Got data-reading error.", http.StatusBadRequest)
		return
	}
	// Check the input data
	// h.l.Printf("Data: %s\n", data)

	fmt.Fprintf(rw, "Hello %s.\n", data)
}
