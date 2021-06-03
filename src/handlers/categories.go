package handlers

import (
	"net/http"
	"text/template"
)

func Categories(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/categories" {
		CodeErreur(w, r, 404, "[Server_Alert] - Error 404 : Page not found")
	}

	files := findPathFiles("./templates/categories.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500, "[Server_Alert] - Error 500 : Template not found -> categories.html")
	}

	ts.Execute(w, nil)
}