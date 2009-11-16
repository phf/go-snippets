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

const timeoutSeconds = 4;

func main() {
	// we consider connection-level errors fatal and panic
	// for those; not suitable for a real web server, but
	// then none of the functions/methods we call on this
	// level *should* fail on a properly configured system
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

		handle_connection(connection);

		error = connection.Close();
		check(error);
	}
}

func check(error os.Error)
{
	if error != nil { panic(error.String()); }
}

func handle_connection(connection io.ReadWriter)
{
	request, path, error := read_request(connection);
	if error == nil {
		send_response(connection, request, "."+path);
	}
}

func read_request(connection io.Reader) (request string, path string, error os.Error)
{
	request = ""; path = "";
	length := 0;

	buffer := make([]byte, 1024);
	length, error = connection.Read(buffer);

	if error != nil { return; }

	// KLUDGE: if we filled the buffer we assume we missed
	// part of the request
	if (length <= 0 || length >= 1024) {
		error = os.NewError("incomplete HTTP request");
		return;
	}

	// TODO: tabs can be separators as well in HTTP,
	// as can multiple spaces; Python's split() is a
	// bit more useful... :-/
	tokens := strings.Split(string(buffer), " ", 0);
	if len(tokens) < 2 {
		error = os.NewError("incomplete HTTP request");
		return;
	}

	request = tokens[0];
	path = tokens[1];
	return;
}

func send_response(connection io.Writer, request string, path string) (error os.Error)
{
	var file *os.File = nil;
	error = nil;

	if strings.ToUpper(request) == "GET" {
		// TODO: use Lstat to check for directory, listing it if
		// possible; currently we just fail for directories...
		file, error = os.Open(path, os.O_RDONLY, 0);
		if error == nil {
			connection.Write(strings.Bytes("HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\n\r\n"));
			io.Copy(connection, file);
			file.Close();
		}
		else {
			connection.Write(strings.Bytes("HTTP/1.0 400 Bad Request\r\nContent-Type: text/html\r\n\r\n<html><h1>400 Bad Request</h1></html>"));
		}
	}
	else {
		connection.Write(strings.Bytes("HTTP/1.0 500 Internal Server Error\r\nContent-Type: text/html\r\n\r\n<html><h1>500 Internal Server Error</h1></html>"));
	}

	return;
}
