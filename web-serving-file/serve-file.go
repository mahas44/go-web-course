package main

import "net/http"

func ServeFile() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogServeFile)
	http.ListenAndServe(":8080", nil)
}

func dogServeFile(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, tobyPic)
}
