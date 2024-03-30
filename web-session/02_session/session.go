package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user Id, user
var dbSessions = map[string]string{} // session id, user id

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV7()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}

	// if the user exist already, get user
	var _user user
	if username, ok := dbSessions[cookie.Value]; ok {
		_user = dbUsers[username]
	}

	// process form submission
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		_user = user{username, firstname, lastname}
		dbSessions[cookie.Value] = username
		dbUsers[username] = _user
	}

	tpl.ExecuteTemplate(w, "index.gohtml", _user)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	username, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	user := dbUsers[username]
	tpl.ExecuteTemplate(w, "bar.gohtml", user)
}
