package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	
	scanner := bufio.NewScanner(conn)
		
	var i int
	var resp_Method, resp_URI string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			the_response := strings.Fields(ln)
			resp_Method = the_response[0]
			resp_URI = the_response[1]
			fmt.Println("METHOD:", resp_Method)
			fmt.Println("URI:", resp_URI)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE HEND OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}

	switch {
	case resp_Method == "GET" && resp_URI == "/":
		handleIndex(conn)
	case resp_Method == "GET" && resp_URI == "/apply":
		handleApply(conn)
	case resp_Method == "POST" && resp_URI == "/apply":
		handleApplyPost(conn)
	default:
		handleDefault(conn)
	}
}

func handleIndex(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="utf-8">
				<title>Index</title>
			</head>
			<body>
				<h1>Our Index</h1>
				<a href="/">index</a><br>
				<a href="/apply">apply</a><br>
			</body>
			</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApply(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="utf-8">
				<title>GET</title>
			</head>
			<body>
				<h1>GET APPLY</h1>
				<a href="/">index</a><br>
				<a href="/apply">apply</a><br>
				<form action="/apply" method="POST">
				<input type="submit" value="submit">
				</form>
			</body>
			</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApplyPost(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="utf-8">
				<title>POST APPLY</title>
			</head>
			<body>
				<h1>POST APPLY</h1>
				<a href="/">index</a><br>
				<a href="/apply">apply</a><br>
			</body>
			</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleDefault(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="utf-8">
				<title>HERE IS DEFAULT</title>
			</head>
			<body>
				<h1>DEFAULT I BE</h1>
			</body>
			</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}