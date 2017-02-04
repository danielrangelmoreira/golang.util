package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	str := []string{"daniel", "rangel", "moreira", "format?=json"}
	q := url.QueryEscape(strings.Join(str, " "))
	fmt.Println(q)
}
