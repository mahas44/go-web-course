package handsonthree

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/*
	We are now going to get "I see you connected" to be written.
	When we used bufio.NewScanner(), our code was reading from an io.Reader that never ended.
	We will now break out of the reading.
	Package bufio has the Scanner type. The Scanner type "provides a convenient interface for reading data".
	When you have a Scanner type, you can call the SCAN method on it.
	Successive calls to the Scan method will step through the tokens (piece of data).
	The default token is a line. The Scanner type also has a TEXT method.
	When you call this method, you will be given the text from the current token.
	Here is how you will use it:
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	Use this code to READ from an incoming connection and print the incoming text to standard out (the terminal).
	When your "ln" line of text is equal to an empty string, break out of the loop.
	Run your code and go to localhost:8080 in your browser.
	What do you find?
*/

func HandsOnThree() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
				break
			}
		}
		fmt.Println("Code got here")
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		io.WriteString(conn, "I see you connected")
		conn.Close()
	}
}
