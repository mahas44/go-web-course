package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var h = "HTTP/1.1 200 OK\r\n"
var cl = "Content-Length: %d\r\n"
var ct = "Content-Type: text/html\r\n"
var head = `<!DOCTYPE html><html lang="en"> <head>	<meta charset="UTF-8"><title>TCP SERVER</title><link rel="icon" href="data:;base64,iVBORw0KGgo=">
</head>	<body>`
var body = `<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	`
var endLine = "</body></html>"

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			mux(conn, ln)
		}
		if ln == "" {
			// header are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {

	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	fmt.Println("***METHOD", method)
	fmt.Println("***URI", uri)

	// multiplexer
	if method == "GET" && uri == "/" {
		index(conn)
	}
	if method == "GET" && uri == "/about" {
		about(conn)
	}
	if method == "GET" && uri == "/contact" {
		contact(conn)
	}
	if method == "GET" && uri == "/apply" {
		apply(conn)
	}
	if method == "POST" && uri == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := head + `<strong>INDEX</strong><br>` + body + endLine
	printConnInfo(conn, body)
}

func about(conn net.Conn) {
	body := head + `<strong>ABOUT</strong><br>` + body + endLine
	printConnInfo(conn, body)
}

func contact(conn net.Conn) {
	body := head + `<strong>CONTACT</strong><br>` + body + endLine
	printConnInfo(conn, body)
}

func apply(conn net.Conn) {
	body := head + `<strong>APPLY</strong><br>` + body +
		`<form method="post" action="/apply">
		<input type="submit value="apply">` + endLine

	printConnInfo(conn, body)
}

func applyProcess(conn net.Conn) {
	body := head + `<strong>APPLY PROCESS</strong><br>` + body + endLine
	printConnInfo(conn, body)
}

func printConnInfo(conn net.Conn, body string) {
	fmt.Fprint(conn, h)
	fmt.Fprintf(conn, cl, len(body))
	fmt.Fprint(conn, ct)
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
