package solutions

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

type hotcat int

func (m hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func One() {
	var d hotdog
	var c hotcat

	// mux := http.NewServeMux()
	http.Handle("/dog/", d) // /dog/something path is work
	http.Handle("/cat", c)  // /cat/something path is not work. Because of / character is not add end of path

	http.ListenAndServe(":8080", nil)
}
