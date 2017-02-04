package main

import "fmt"

// func main() {
// 	demPanic()
// 	fmt.Println("This was written after panic")
// }

// func demPanic(){
// 	defer func() {
// 		fmt.Println(recover())
// 	}()
// 	panic("PANIC")
// }

func main() {
	fmt.Println(safeDiv(4, 0))
	fmt.Println(safeDiv(4, 2))
}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())

	}()

	return num1 / num2
}
