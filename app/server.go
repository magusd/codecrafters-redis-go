package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func Ping(conn net.Conn) {
	pong := []byte("+PONG\r\n")
	_, err := conn.Write(pong)
	if err != nil {
		log.Fatal("Failed to respond to ping")
		fmt.Println(err)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer listener.Close()

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buffer := make([]byte, 1024)

	for {
		cmd, _ := connection.Read(buffer)
		fmt.Println("Received: ", string(cmd))
		Ping(connection)
	}

	fmt.Print("\nStopping the server...")

}
