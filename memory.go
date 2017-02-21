package main

import (
	"fmt"
	"time"
)

var (
	chx = make(chan int)
	chy = make(chan int)
)

func main() {
	go concurrent()

	for i := 0; i < 10; i++ {
		go func() {
			chx <- 1
			fmt.Print("y:", <-chy, " ")
			fmt.Println()

		}()
		go func() {
			chy <- 1
			fmt.Print("x:", <-chx, " ")
			fmt.Println()

		}()

	}

	time.Sleep(1 * time.Second)
}

func concurrent() {
	var x, y int
	for {
		select {
		case x = <-chx:
		case y = <-chy:
		case chx <- x:
		case chy <- y:
		}
	}

}
