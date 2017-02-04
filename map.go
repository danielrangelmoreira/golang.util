package main

import (
	"fmt"
	"sort"
	"strings"
)

func printAges(hash map[string]int, keys []string) {
	once := true
	for _, name := range keys {

		if once == true {
			fmt.Printf("Name\tValue\n")
		}
		fmt.Printf("%s\t%d\n", strings.Title(name), hash[name])
		once = false
	}
}

func main() {
	var ages = make(map[string]int)
	var names []string

	ages["daniel"] = 34
	ages["lorena"] = 31
	ages["talita"] = 30
	ages["avanir"] = 55
	ages["jose"] = 40
	ages["heitor"] = 23
	ages["marcos"] = 19
	ages["joaquim"] = 10

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	printAges(ages, names)

}
