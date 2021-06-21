package handlers

import (
	"html/template"
	"net/http"
	//"text/template"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/posts.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, nil)

}
