/*
	A web server inspired by my friend Raluca. Her's is
	128 bytes of Python. The Go version is a bit longer
	but who cares. :-D There are two versions, this one
	tries to be as short as possible (modulo formatting).

	http://ralucam.tumblr.com/post/178403091/twitcode

	Here's a version of Raluca's code that's compatible
	with my Go implementation:

	import socket
	s=socket.socket();s.bind(('',8080));s.listen(9)
	while(1):c,a=s.accept();c.send(open("."+c.recv(1024).split()
	[1]).read());c.close()
*/

package main

import "net"
import "strings"
import "os"
import "io"

func main() {
	listener, _ := net.Listen("tcp", ":8080")

	for {
		connection, _ := listener.Accept()

		buffer := make([]byte, 1024)
		connection.Read(buffer)

		tokens := strings.Split(string(buffer), " ", 0)

		file, _ := os.Open("."+tokens[1], os.O_RDONLY, 0)
		io.Copy(connection, file)
		file.Close()

		connection.Close()
	}
}
