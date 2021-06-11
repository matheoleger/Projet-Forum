package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func CodeErreur(w http.ResponseWriter, r *http.Request, status int) {

	files := findPathFiles("./templates/error.html")

	t, err := template.ParseFiles(files...)

	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, ErrorType(w, r, status))

}

func ErrorType(w http.ResponseWriter, t *http.Request, status int) string {
	var errorstr string
	if status == 400 {
		errorstr = `Error 400 : Bad Request `
	}
	if status == 404 {
		errorstr = `Error 404 : Page Not Found`
	}
	if status == 500 {
		errorstr = `Error 500 : Internal Server Error`
	}
	return errorstr
}
