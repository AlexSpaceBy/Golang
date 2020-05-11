// This is a simple client to test our server

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	// We dial to the server by using local address
	conn, ok := net.Dial("tcp", "127.0.0.1:8080")
	if ok != nil {
		fmt.Println(ok)
		return
	}
	defer conn.Close()

	// We print received JSON message directly to console
	io.Copy(os.Stdout, conn)

}
