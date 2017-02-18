package main

import (
	"fmt"
	"sync"
)

var letters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	var wg sync.WaitGroup
	//var done = make(chan struct{})
	wg.Add(2)

	go func() {
		for _, number := range numbers {
			fmt.Println(number)
		}
		//done <- struct{}{}
		wg.Done()
	}()

	go func() {
		for _, letter := range letters {
			fmt.Printf("%q\n", letter)
		}
		//done <- struct{}{}
		wg.Done()

	}()

	//for i := 0; i < 2; i++ {
	//<-done

	//}
	wg.Wait()
}
