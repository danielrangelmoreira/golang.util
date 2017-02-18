package main

import (
	"flag"

	"io"
	"log"
	"net"
	"os"
)

func main() {
	var port = flag.String("port", "8000", "especifique uma porta")
	var host = flag.String("host", "localhost", "especifique um endereço")
	flag.Parse()

	conn, err := net.Dial("tcp", *host+":"+*port)

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
