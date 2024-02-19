package formfile

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// ex.2
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("04_form-file/templates/*"))
}

func Foo(w http.ResponseWriter, req *http.Request) {
	var str string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		// open
		file, fileHeader, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		// for your infomation
		fmt.Println("\nfile:", file, "\nfile header:", fileHeader, "\nerr", err)

		// read
		bs, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		str = string(bs)

		// ex.2
		// store on server
		destination, err := os.Create(filepath.Join("04_form-file/user/", fileHeader.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer destination.Close()

		_, err = destination.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// ex.1
	// io.WriteString(w, `
	// 	<form method="POST" enctype="multipart/form-data">
	// 	<input type="file" name="q">
	// 	<input type"submit">
	// 	</form>
	// 	<br>`+str)

	// ex.2
	tpl.ExecuteTemplate(w, "index.gohtml", str)
}
