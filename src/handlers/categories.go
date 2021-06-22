package handlers

import (
	"net/http"
	"text/template"

	bdd "../database"
)

func Categories(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/categories" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/categories.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

<<<<<<< HEAD
	

	ts.Execute(w, nil)
=======
	page := bdd.Page{Categories: bdd.GetCategory()}

	ts.Execute(w, page)
>>>>>>> b7bafd2922cccab9f7b503849373295f1b1e37c4
}
