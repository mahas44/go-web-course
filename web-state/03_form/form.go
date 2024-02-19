package form

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

func Foo(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
