/*
	A web server inspired by my friend Raluca. Her's is
	128 bytes of Python. The Go version is a bit longer
	but who cares. There are two versions, this one tries
	to be as short as possible (modulo formatting).
*/

package main

import (
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")

	for {
		connection, _ := listener.Accept()

		buffer := make([]byte, 1024)
		connection.Read(buffer)

		tokens := strings.Split(string(buffer), " ")

		file, _ := os.Open("." + tokens[1])
		io.Copy(connection, file)
		file.Close()

		connection.Close()
	}
}
