// This is a simple client package

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	for {
		// We read from the console
		fmt.Println("Enter the message:")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// If exit - the program exits
		if strings.TrimSpace(text) == "exit" {
			os.Exit(0)
		}

		// We dial to the server
		conn, ok := net.Dial("tcp", "127.0.0.1:8083")
		if ok != nil {
			fmt.Println(ok)
			return
		}

		// We send the message to the server
		conn.Write([]byte(text))

		// We read the message assuming that message is less than 100b
		buffer := make([]byte, 100)
		n, ok := conn.Read(buffer)
		if ok != nil {
			fmt.Println(ok)
			return
		}

		fmt.Println("\nThe answer from server:")

		// We convert message to string by trimming all unnecessary data
		slice := buffer[:n]
		message := string(slice)

		fmt.Println(message)

		conn.Close()
	}
}
