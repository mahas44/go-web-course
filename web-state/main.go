package main

import (
	"html/template"
	"log"
	"net/http"
)

// ex.3
var tpl *template.Template

// ex.3
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// ex.3
type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// ex.1
// passing values through the URL
// func foo(w http.ResponseWriter, req *http.Request) {
// 	v := req.FormValue("q")
// 	io.WriteString(w, "Do my search: "+v)
// }

// ex.2
// passing data with form
// func foo(w http.ResponseWriter, req *http.Request) {
// 	v := req.FormValue("q")
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	io.WriteString(w, `
// 		<form method="post">
// 			<input type="text" name="q">
// 			<input type="submit">
// 		</form>
// 		<br>`+v)
// }

// ex.3
func foo(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
