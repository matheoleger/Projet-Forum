package handlers

import (
	"net/http"
	"strconv"
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

	wichpage := r.URL.Query().Get("page")
	wichpageInt, _ := strconv.Atoi(wichpage)

	perpage := r.URL.Query().Get("perpage")
	perpageInt, _ := strconv.Atoi(perpage)

	page := bdd.Page{Categories: bdd.GetCategory(perpageInt, wichpageInt)}

	ts.Execute(w, page)
}
