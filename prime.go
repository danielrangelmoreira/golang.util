package main

import (
	"fmt"
)

func isprime(num int) bool {
	if num%2 == 0 {
		return false
	}
	for i := 3; i < num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 0; i < 31; i++ {
		if isprime(i) {
			fmt.Println(i)
		}
	}

}
