package main

import (
	"fmt"
)

// [5, 4, 3, 2, 1, 0]
// [0, 1, 2, 3, 4, 5]
// [2, 3, 4, 5, 0, 1]
func rotateN(ints []int, n int) {
	for i, j := 0, len(ints)-(n+1); i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

func ptrReverse(ints *[]int) {
	for i, j := 0, len(*ints)-1; i < j; i, j = i+1, j-1 {
		(*ints)[i], (*ints)[j] = (*ints)[j], (*ints)[i]
	}
}

func rInPlace(ints []int) {
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

// [0, 1, 2, 3, 4, 5]
func reverse(ints []int) []int {
	var out []int
	for i := len(ints) - 1; i >= 0; i-- {
		out = append(out, ints[i])
	}
	return out
}

func main() {
	ints := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(ints)
	fmt.Println(reverse(ints))
	rInPlace(ints)
	fmt.Printf("rInPlace: %v\n", ints)
	rInPlace(ints)
	fmt.Printf("rInPlace: %v\n", ints)

	ptrReverse(&ints)
	fmt.Printf("rInPlace: %v\n", ints)
	rotateN(ints, 2)
	fmt.Printf("rotate: %v\n", ints)
}
