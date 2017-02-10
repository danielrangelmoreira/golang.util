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

const firstmsg = `/fs/fs3_sys#Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36$(</fs/fs3_service(</fs/fs3_u_1_1(<`

var addr = flag.String("addr", "p8.fsdatacentre.com", "http service address")

var myDialer = &websocket.Dialer{
	Proxy: http.ProxyFromEnvironment,
	//EnableCompression: true,
}

func ReadSocket(c *websocket.Conn) {
	log.Println("Listening...")

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	myHeader := make(http.Header, 1)

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

	scanner := bufio.NewScanner(os.Stdin)

	err = c.WriteMessage(websocket.BinaryMessage, []byte(firstmsg))
	if err != nil {
		log.Println("write:", err)
		return
	}

	go ReadSocket(c)

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
