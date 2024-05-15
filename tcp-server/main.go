package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func Listen(network, address string) (net.Listener, error) {
	return net.Listen(network, address)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	address := fmt.Sprintf("0.0.0.0:%s", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Listening on " + address)
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleIncomingRequest(connection)
	}

}

func handleIncomingRequest(conn net.Conn) {
	defer conn.Close()
	// https://www.rfc-editor.org/rfc/rfc9112.html#section-2.1

	msg := []byte("Hello World!")

	// status-line
	// https://www.rfc-editor.org/rfc/rfc9112.html#section-4
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))

	// https://www.rfc-editor.org/rfc/rfc9112.html#section-6.3-3
	conn.Write([]byte(fmt.Sprintf("content-length: %d\r\n", len(msg))))

	// end of header fields
	// https://www.rfc-editor.org/rfc/rfc9112.html#section-2.1
	conn.Write([]byte("\r\n"))
	// message body
	// https://www.rfc-editor.org/rfc/rfc9112.html#section-2.1
	conn.Write(msg)
}
