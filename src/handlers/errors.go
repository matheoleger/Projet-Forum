package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func CodeErreur(w http.ResponseWriter, r *http.Request, status int) {
	//var errorType string
	const colorRed = "\033[31m"

	files := findPathFiles("./templates/error.html")

	t, err := template.ParseFiles(files...)

	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, ErrorType(w, r, t, status))

}

func ErrorType(w http.ResponseWriter, t *http.Request, err *template.Template, status int) string {
	var errorstr string
	if status == 400 {
		errorstr = `Error 400 : Bad Request `
		//err.Execute(w, `<h1> Error 400 : Bad Request </h1>`)
	}
	if status == 404 {
		errorstr = `Error 404 : Page Not Found`
		//err.Execute(w, `<h1> Error 404 : Page Not Found </h1>`)
	}
	if status == 500 {
		errorstr = `Error 500 : Internal Server Error`
		//err.Execute(w, `<h1> Error 500 : Internal Server Error </h1>`)
	}
	return errorstr
}
