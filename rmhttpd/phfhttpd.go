/*
	A web server inspired by my friend Raluca. Her's is
	128 bytes of Python. The Go version is a bit longer
	but who cares. :-D There are two versions, this one
	tries to do proper error checking and implements a
	minimal HTTP protocol as well.

	http://ralucam.tumblr.com/post/178403091/twitcode
*/

package main

import "net"
import "strings"
import "os"
import "io"

const timeoutSeconds = 60;

func check(error os.Error)
{
	if error != nil { panic(error.String()); }
}

func main() {
	listener, error := net.Listen("tcp", "localhost:8080");
	check(error);

	for
		connection, error := listener.Accept();
		error == nil;
		connection, error = listener.Accept()
	{
		check(error);

		error := connection.SetTimeout(timeoutSeconds*1000*1000*1000);
		check(error);

		buffer := make([]byte, 1024);
		_, error = connection.Read(buffer);
		if error == nil {
			tokens := strings.Split(string(buffer), " ", 0);
			if strings.ToUpper(tokens[0]) == "GET" {
				name := "." + tokens[1];
				file, error := os.Open(name, os.O_RDONLY, 0);
				if error == nil {
					connection.Write(strings.Bytes("HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\n\r\n"));
					io.Copy(connection, file);
					file.Close();
				}
				else {
					connection.Write(strings.Bytes("HTTP/1.0 400 Bad Request\r\n\r\n"));
				}
			}
		}

		error = connection.Close();
		check(error);
	}
}
