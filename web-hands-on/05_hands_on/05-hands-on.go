package handsonfive

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/*
	Before we WRITE our RESPONSE, let's WRITE to our RESPONSE the STATUS LINE and some RESPONSE HEADERS.
	Remember the request line and status line:
	
	REQUEST LINE 
	GET / HTTP/1.1 
	method SP request-target SP HTTP-version CRLF
	https://tools.ietf.org/html/rfc7230#section-3.1.1

	RESPONSE (STATUS) LINE
	HTTP/1.1 302 Found 
	HTTP-version SP status-code SP reason-phrase CRLF 
	https://tools.ietf.org/html/rfc7230#section-3.1.2

	Write the following strings to the response - use io.WriteString for all of the following except the second and third:
	"HTTP/1.1 200 OK\r\n"
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	"\r\n"
	Look in your browser "developer tools" under the network tab.
	Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.

*/

func HandsOnFive() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
