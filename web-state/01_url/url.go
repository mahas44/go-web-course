package formurl

import (
	"io"
	"net/http"
)

// passing values through the URL
func Foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}
