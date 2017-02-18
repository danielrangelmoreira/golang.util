package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range d {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := d[item]
		if !ok {
			err := fmt.Sprintf("item not found: %s\n", item)
			http.Error(w, err, http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", item, price)
	default:
		err := fmt.Sprintf("page not found: %s\n", r.URL)
		http.Error(w, err, http.StatusNotFound)
	}
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
