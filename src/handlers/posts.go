package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	bdd "../database"
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

	// item := GetPost()

	categoryName := r.URL.Query().Get("category")

	fmt.Println(categoryName)

	page := bdd.Page{Posts: bdd.GetPostByCategory(categoryName)}

	// ts.Execute(w, item)
	ts.Execute(w, page)

}

}
