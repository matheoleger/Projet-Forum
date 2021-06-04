package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		CodeErreur(w, r, 404)
	}

	files := findPathFiles("./templates/home.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
	}

	ts.Execute(w, nil)
}
