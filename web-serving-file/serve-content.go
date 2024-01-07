package main

import (
	"net/http"
	"os"
)

func ServeContent() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPicServeContent)
	http.ListenAndServe(":8080", nil)
}

func dogPicServeContent(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open(tobyPic)
	if err != nil {
		http.Error(w, "file not fount", 404)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
