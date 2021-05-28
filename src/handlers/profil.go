package handlers

import (
	"net/http"
	"text/template"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profil" {
		CodeErreur(w, r, 404, "Page not found")
	}

	files := findPathFiles("./templates/profil.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500, "Template not found : profil.html")
	}

	ts.Execute(w, nil)
}
