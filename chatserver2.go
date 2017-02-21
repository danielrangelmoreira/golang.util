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
	select {
	case <-timer.C:
		conn.Close()
	}
}
func getScreenName(conn net.Conn) (name string) {
	fmt.Fprintln(conn, "Please input your name:")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		name = input.Text()
		break
	}
	return name

}
func handleConn(conn net.Conn) {
	ch := make(chan string, 10)
	who := getScreenName(conn) //conn.RemoteAddr().String()
	cli := client{ch, who}
	timer := time.NewTimer(5 * time.Minute)

	go clientWrite(conn, ch)
	go idleCloser(conn, timer)

	ch <- "You are: " + who + " at " + conn.RemoteAddr().String()
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		if !timer.Stop() {
			<-timer.C
		}
		timer.Reset(5 * time.Minute)
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
				select {
				case cli.ch <- msg:
				default:
				}
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
