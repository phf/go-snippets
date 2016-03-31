/*
	A web server inspired by my friend Raluca. Her's is
	128 bytes of Python. The Go version is a bit longer
	but who cares. There are two versions, this one tries
	to do some error checking, supports directory listings,
	and implements a minimal HTTP protocol as well.
*/

package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	timeout                  = time.Second * 4
	folkloreMaxRequestLength = 8192
)

// Abort with a panic if there was an error. Not exactly suitable for a
// production server, is it? Well, we're mostly interested in concision,
// not reliability.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Read an HTTP request. The returned path is not sanitized.
func readRequest(conn io.Reader) (req string, path string, err error) {
	buffer := make([]byte, folkloreMaxRequestLength)
	_, err = conn.Read(buffer)
	if err != nil {
		return "", "", err
	}

	tokens := strings.Fields(string(buffer))
	if len(tokens) < 2 {
		err = errors.New("incomplete HTTP request")
		return "", "", err
	}

	req = tokens[0]
	path = tokens[1]
	return req, path, nil
}

// Send an HTTP response. The path is supposed to be sanitized already.
func sendResponse(conn io.Writer, req string, path string) (err error) {
	var file *os.File
	var info os.FileInfo

	if strings.ToUpper(req) != "GET" {
		conn.Write([]byte("HTTP/1.0 501 Not Implemented\r\nContent-Type: text/html\r\n\r\n<html><h1>501 Not Implemented</h1></html>"))
		err = errors.New("only GET is implemented")
		return err
	}

	info, err = os.Lstat(path)
	if err != nil {
		conn.Write([]byte("HTTP/1.0 404 Not Found\r\nContent-Type: text/html\r\n\r\n<html><h1>404 Not Found</h1></html>"))
		return err
	}

	if !info.Mode().IsRegular() && !info.Mode().IsDir() {
		conn.Write([]byte("HTTP/1.0 400 Bad Request\r\nContent-Type: text/html\r\n\r\n<html><h1>400 Bad Request</h1></html>"))
		err = errors.New("not a regular file or directory")
		return err
	}

	if info.Mode().IsDir() {
		file, _ = os.Open(path)           // TODO error handling
		names, _ := file.Readdirnames(-1) // TODO error handling
		conn.Write([]byte("HTTP/1.0 200 OK\r\nContent-Type: text/html\r\n\r\n<html><h1>Directory Listing</h1>\n"))
		for _, name := range names {
			conn.Write([]byte(fmt.Sprintf("<a href=\"%s\">%s</a><br/>\n", name, name)))
		}
		conn.Write([]byte("</html>"))
		return nil
	}

	file, _ = os.Open(path) // TODO error handling
	conn.Write([]byte("HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\n\r\n"))
	io.Copy(conn, file)
	file.Close()
	return nil
}

// Handle a connection by reading the request and sending the response.
func handleConnection(conn io.ReadWriter) {
	req, path, err := readRequest(conn)
	path = filepath.Clean(path)
	if err == nil {
		sendResponse(conn, req, "./"+path)
	}
}

// Run the web server.
func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	check(err)

	for {
		conn, err := listener.Accept()
		check(err)

		err = conn.SetDeadline(time.Now().Add(timeout))
		check(err)

		handleConnection(conn)

		err = conn.Close()
		check(err)
	}
}
