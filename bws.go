package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

const firstmsg = `#P__time,S_652870591E03400E87B6CF31670630CE000003`

var addr = flag.String("addr", "premws-pt3.365lpodds.com", "http service address")

var myDialer = &websocket.Dialer{
	Proxy:             http.ProxyFromEnvironment,
	EnableCompression: true,
}
var myHeader = http.Header{
	"Origin":                 []string{"https://www.bet365.com"},
	"Sec-WebSocket-Protocol": []string{"zap-protocol-v1"},
}

func ReadSocket(c *websocket.Conn) {
	log.Println("Listening...")

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s\n", message)
	}
}

func main() {
	flag.Parse()

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/zap"}
	q := u.Query()
	q.Set("uid", "8826877042224068")
	u.RawQuery = q.Encode()

	log.Printf("connecting to %s", u.String())

	c, resp, err := myDialer.Dial(u.String(), myHeader)
	if err != nil {
		log.Fatal("dial:", err)

	} else if resp.StatusCode == http.StatusSwitchingProtocols {
		log.Printf("connected! \n")

	}
	defer c.Close()

	err = c.WriteMessage(websocket.BinaryMessage, []byte(firstmsg))
	if err != nil {
		log.Println("write:", err)
		return
	}

	go ReadSocket(c)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Bytes()
		err := c.WriteMessage(websocket.BinaryMessage, msg)
		log.Println("wrote")
		if err != nil {
			log.Println("write:", err)
			return
		}

	}
}
