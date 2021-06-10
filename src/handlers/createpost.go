package handlers

import (
	"net/http"
	"text/template"
)

func Creationpost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/creationpost" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/createpost.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, nil)
}
