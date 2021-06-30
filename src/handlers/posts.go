package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	bdd "../database"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts" {
		fmt.Println("ici cest la merde ")
		CodeErreur(w, r, 404)
		return
	}

	wichpage := r.URL.Query().Get("page")
	wichpageInt, _ := strconv.Atoi(wichpage)

	perpage := r.URL.Query().Get("perpage")
	perpageInt, _ := strconv.Atoi(perpage)

	categoryName := r.URL.Query().Get("category")

	fmt.Println("ici cest le params : " + categoryName)

	category := bdd.Category{Name: categoryName}
	var categories []bdd.Category
	categories = append(categories, category)

	posts := bdd.GetPostByCategory(categoryName, perpageInt, wichpageInt)

	if VerifyCookie(w, r) {
		username := bdd.GetProfil(w, r).Username

		for i, postEl := range posts {
			postEl.LikeInfo = bdd.IsLiked("post", username, postEl.Id_post)

			posts[i] = postEl
		}
	}

	page := bdd.Page{Categories: categories, Posts: posts}

	files := findPathFiles("./templates/posts.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("ici cest la merde ")
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, page)

}

func PostsContent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts/content" {
		fmt.Println("ici cest la merde content")
		CodeErreur(w, r, 404)
		return
	}

	postname := r.URL.Query().Get("post")
	postnameint, _ := strconv.Atoi(postname)

	wichpage := r.URL.Query().Get("page")
	wichpageInt, _ := strconv.Atoi(wichpage)

	perpage := r.URL.Query().Get("perpage")
	perpageInt, _ := strconv.Atoi(perpage)

	db := OpenDataBase()
	post := bdd.GetPost(db, postnameint)
	db.Close()
	var posts []bdd.Post
	posts = append(posts, post)

	comments := bdd.GetComments(postnameint, perpageInt, wichpageInt*perpageInt)

	if VerifyCookie(w, r) {
		username := bdd.GetProfil(w, r).Username

		for i, commentEl := range comments {
			commentEl.LikeInfo = bdd.IsLiked("comment", username, commentEl.Id_comment)

			comments[i] = commentEl
		}
	}

	page := bdd.Page{Posts: posts, Comments: comments}

	files := findPathFiles("./templates/post_content.html")
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("ici cest la merde dans le parsefile")
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, page)

}
