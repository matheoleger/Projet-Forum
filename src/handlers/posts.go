package handlers

import (
	"net/http"
	"text/template"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts" {
		CodeErreur(w, r, 404, "Page not found")
	}

	files := findPathFiles("./templates/posts.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500, "Template not found : posts.html")
	}

	ts.Execute(w, nil)
}
