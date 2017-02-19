package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	ch   chan string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}
func idleCloser(conn net.Conn, timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			conn.Close()

		}
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)
	who := conn.RemoteAddr().String()
	cli := client{ch, who}
	timer := time.NewTimer(60 * time.Second)

	go clientWrite(conn, ch)
	go idleCloser(conn, timer)
	ch <- "You are: " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(60 * time.Second)
	}

	leaving <- cli
	messages <- who + " has left!"
	conn.Close()
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case new := <-entering:
			clients[new] = true
			for cli := range clients {
				new.ch <- cli.name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}

}
