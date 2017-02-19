package main

import (
	"fmt"
	"time"
)

const (
	timeout = 10 * time.Second
)

func main() {
	timer := time.NewTimer(timeout)
	tic := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			goprint()
		case <-tic.C:
			tictac()
		}
	}

}
func tictac() {
	fmt.Println("tic")
}

func goprint() {
	fmt.Println("TIMEOUT!")

}
