package main

import (
	"bufio"
	"fmt"
	"os"
)

type chholder chan string

var (
	register   = make(chan chholder)
	unregister = make(chan chholder)
	send       = make(chan string)
)

func main() {
	go handleChan()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var ch = make(chan string)

		register <- ch
		go PrintOut(ch)

		msg := scanner.Text()
		if msg == "exit" {
			break
		}
		send <- msg

	}
}
func PrintOut(ch chan string) {
	for msg := range ch {
		fmt.Fprintln(os.Stdout, msg)
	}
}

func handleChan() {
	channel := make(map[chholder]bool)

	for {
		select {
		case ch := <-register:
			channel[ch] = true

		case ch := <-unregister:
			delete(channel, ch)
			close(ch)

		case msg := <-send:
			for ch := range channel {
				ch <- msg
			}
		}

	}

}
