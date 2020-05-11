// This is a server that for every request sends answer in JSON format

package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// Struct for header to match the homework
type header struct {
	Accept    [1]string
	UserAgent [1]string
}

// request struct for server to answer in JSON like format
type request struct {
	Host        string
	User_agent  string
	Request_uri string
	Headers     header
}

func main() {

	// To show that server is running
	fmt.Println("Turning on the server")

	// We create the response message from the server
	message := &request{
		Host:        "127.0.0.1:8080",
		User_agent:  "curl/7.67.0",
		Request_uri: "www.google.com",
		Headers: header{
			Accept:    [1]string{"*/*"},
			UserAgent: [1]string{"curl/7.67.0"},
		},
	}

	// We listen to 8080 port by using tcp protocol
	listener, ok := net.Listen("tcp", ":8080")
	if ok != nil {
		fmt.Println(ok)
		return
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
			return
		}

		// We convert server response to the JSON
		jsonMessage, ok := json.Marshal(message)
		if ok != nil {
			fmt.Println(ok)
			return
		}

		// We answer the response
		conn.Write(jsonMessage)

		fmt.Println("Message is sent")

		conn.Close()
	}
}
