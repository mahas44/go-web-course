package main

import (
	"io"
	"net/http"
	"os"
)

func IOCopy() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open(tobyPic)
	if err != nil {
		http.Error(w, "file not fount", 404)
	}

	defer f.Close()
	io.Copy(w, f)
}
