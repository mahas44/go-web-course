package handsonfour

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/*
	Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".
	Pass the connection of type net.Conn as an argument into this function.
	Add "go" in front of the call to "serve" to enable concurrency and multiple connections.
*/

func HandsOnFour() {
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
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, "Here we WRITE to the response")
}
