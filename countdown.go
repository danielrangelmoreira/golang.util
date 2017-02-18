package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	tick := time.Tick(1 * time.Second)
	fmt.Println("Comencing countdown!")

	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			abort <- struct{}{} //abort anyway
		}
		abort <- struct{}{}
	}()

	for count := 10; count > 0; count-- {
		fmt.Println(count)
		select {
		case <-tick:
		//do nothing
		case <-abort:
			fmt.Println("Launch Aborted!")
			return
		}
	}
	launch()
}
func launch() {
	fmt.Println("Launch!")
}
