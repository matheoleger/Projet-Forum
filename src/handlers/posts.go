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
		fmt.Println("ici cest la merde ")
		CodeErreur(w, r, 404)
		return
	}

	categoryName := r.URL.Query().Get("category")

	fmt.Println("ici cest le params : " + categoryName)

	category := bdd.Category{Name: categoryName}
	var categories []bdd.Category
	categories = append(categories, category)

	page := bdd.Page{Categories: categories, Posts: bdd.GetPostByCategory(categoryName)}

	files := findPathFiles("./templates/posts.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("ici cest la merde ")
		CodeErreur(w, r, 500)
		return
	}

	// item := GetPost()

	// ts.Execute(w, item)
	ts.Execute(w, page)

}
