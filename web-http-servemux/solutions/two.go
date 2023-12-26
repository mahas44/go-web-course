package solutions

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func Two() {
	http.HandleFunc("/dog/", d) // /dog/something path is work
	http.HandleFunc("/cat", c)  // /cat/something path is not work. Because of / character is not add

	http.ListenAndServe(":8080", nil)
}
