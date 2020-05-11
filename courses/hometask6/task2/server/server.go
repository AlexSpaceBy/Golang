// This is a server

package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {

	// To show that server is running
	fmt.Println("Turning on the server")

	// We are listening to 8080 port by using tcp protocol
	listener, ok := net.Listen("tcp", ":8083")
	if ok != nil {
		fmt.Println(ok)
	}
	defer listener.Close()

	// If everything is ok the server is ready for handling the requests
	fmt.Println("The server is ready for receiving")

	// Infinite loop for requests, to stop the server CTRL+C
	for {

		// We accept every request
		conn, ok := listener.Accept()
		if ok != nil {
			fmt.Println(ok)
			conn.Close()
			continue
		}

		// We read the message assuming that message is less than 100b
		buffer := make([]byte, 100)
		n, ok := conn.Read(buffer)
		if ok != nil {
			fmt.Println(ok)
			return
		}

		// We convert message to string by trimming all unnecessary data
		slice := buffer[:n]
		message := string(slice)

		// We delete all control characters
		text := strings.TrimSpace(message)

		// We test input for int
		testInt, ok := strconv.Atoi(text)
		if ok != nil {
			conn.Write([]byte(strings.ToUpper(message)))
		} else {
			conn.Write([]byte(strconv.Itoa(testInt * 2)))
		}

		conn.Close()
	}
}
