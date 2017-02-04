package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	//"strings"
)

func main() {
	//var sz = []byte(`{"daniel":"rangel"}`)
	for _, ur := range os.Args[1:] {
		resp, err := http.PostForm(ur, url.Values{
			"Nome":  {"Daniel"},
			"Idade": {"35"},
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch GET: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch ReadAll: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)

	}
}
