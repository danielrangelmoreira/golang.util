package main

import (
	"flag"

	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	var port = flag.Int("port", 8000, "especifique uma porta")
	var host = flag.String("host", "localhost", "especifique um endereço")
	flag.Parse()

	//fmt.Println(*host + ":" + strconv.Itoa(*port))
	conn, err := net.Dial("tcp", *host+":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

}
