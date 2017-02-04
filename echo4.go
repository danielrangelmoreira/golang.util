package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing new line")
var sep = flag.String("s", " ", "insert separator")

func main() {
	flag.Parse()
  for _, iterator := range flag.Args(){ 
    fmt.Printf("%+v\n",iterator) 
  }
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Print("\n")
	}
}
