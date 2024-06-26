package parseform

import (
	"log"
	"net/http"
	"text/template"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("parseform/index.gohtml"))
}

func ParseForm() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
