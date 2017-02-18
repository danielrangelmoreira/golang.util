package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "p1.fsdatacentre.com", "http service address")

var myDialer = &websocket.Dialer{
	Proxy:             http.ProxyFromEnvironment,
	EnableCompression: true,
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	myHeader := make(http.Header, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/WebSocketConnection-Secure"}

	myHeader["Origin"] = []string{"http://www.resultados.com"}
	myHeader["Sec-WebSocket-Protocol"] = []string{"zap-protocol-v1"}

	log.Printf("connecting to %s", u.String())

	c, resp, err := myDialer.Dial(u.String(), myHeader)

	if err != nil {
		log.Fatal("dial:", err)
	} else if resp.StatusCode == http.StatusSwitchingProtocols {
		log.Printf("connected! \n")
	}
	defer c.Close()

	done := make(chan struct{})
	count := make(chan int, 3)
	go func() {
		defer close(done)

		for {
			log.Println("im back")
			log.Println(<-count)
			_, message, err := c.NextReader()

			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	i := 0
	for {

		select {
		case <-ticker.C:
			log.Println("writing")
			err := c.WriteMessage(websocket.BinaryMessage, []byte("4`"))
			log.Println("wrote")
			if err != nil {
				log.Println("write:", err)
				return
			}

			count <- i
			i++
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bye"))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}
