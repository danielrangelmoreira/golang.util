package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buff bytes.Buffer
	var str = "Hello World!"

	buff.WriteString(str)
	slice1 := buff.Next(5)
	slice2 := buff.Next(10)

	fmt.Printf("1: %s / 2: %s \n", slice1, slice2)
}
