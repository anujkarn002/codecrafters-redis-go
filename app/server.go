package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

const (
	MAXBUFFER = 5000
	PONG      = "+PONG\r\n"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Server Connected")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Received %v\n", string(buffer[0:n]))
		handlePing(conn)
	}
}

func handlePing(conn net.Conn) string {
	conn.Write([]byte(PONG))
	return PONG
}
