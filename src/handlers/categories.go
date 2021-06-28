package handlers

import (
	"fmt"
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

	page := bdd.Page{Categories: bdd.GetCategory()}

	ts.Execute(w, page)
}

func insertBridge(B_id_post int, B_id_category string) {
	db := OpenDataBase()

	statement, err := db.Prepare("INSERT INTO bridge (B_id_post, B_id_category) VAlUES (?, ?)")

	if err != nil {
		fmt.Println("error prepare ")
		return
	}
	statement.Exec(B_id_post, B_id_category)

	defer db.Close()
}
