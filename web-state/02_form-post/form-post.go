package formpost

import (
	"io"
	"net/http"
)

// passing data with form
func Foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post">
			<input type="text" name="q">
			<input type="submit">
		</form>
		<br>`+v)
}
