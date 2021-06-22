package handlers

import (
	"net/http"
	"text/template"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profil" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/profil.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, nil)
}
