package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

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

func PostsContent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts/content" {
		fmt.Println("ici cest la merde content")
		CodeErreur(w, r, 404)
		return
	}

	postname := r.URL.Query().Get("post") //categoryName := r.URL.Query().Get("category")
	postnameint, _ := strconv.Atoi(postname)

	fmt.Println("ici cest le params : " + postname) //fmt.Println("ici cest le params : " + categoryName)

	db := OpenDataBase()
	post := bdd.GetPost(db, postnameint)
	db.Close()
	// post := bdd.Post{Id_post: postnameint} //category := bdd.Category{Name: categoryName}
	var posts []bdd.Post        //var categories []bdd.Category
	posts = append(posts, post) //categories = append(categories, category)

	page := bdd.Page{Posts: posts, Comments: bdd.GetComments(postnameint)} //page := bdd.Page{Categories: categories, Posts: bdd.GetPostByCategory(categoryName)}
	fmt.Print(page)

	files := findPathFiles("./templates/post_content.html")
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("ici cest la merde dans le parsefile")
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, page) //page

}
