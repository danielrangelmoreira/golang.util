package main

import (
	"fmt"
)

func main() {
	var a, b, c int64
	b = 3
	c = 4
	fmt.Println(a, b, c)
	a = b + c
	fmt.Println(a, b, c)
}
