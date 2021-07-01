package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	bdd "../database"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	// Gestion d'erreur 404
	if r.URL.Path != "/posts" {
		CodeErreur(w, r, 404)
		return
	}

	// Récupération de certains éléments de l'URL
	wichpage := r.URL.Query().Get("page")
	wichpageInt, _ := strconv.Atoi(wichpage)

	perpage := r.URL.Query().Get("perpage")
	perpageInt, _ := strconv.Atoi(perpage)

	categoryName := r.URL.Query().Get("category")

	category := bdd.Category{Name: categoryName}
	var categories []bdd.Category
	categories = append(categories, category)

	posts := bdd.GetPostByCategory(categoryName, perpageInt, wichpageInt)

	// Vérification de la connexion de l'utilisateur
	if VerifyCookie(w, r) {
		username := bdd.GetProfil(w, r).Username

		for i, postEl := range posts {
			postEl.LikeInfo = bdd.IsLiked("post", username, postEl.Id_post)

			posts[i] = postEl
		}
	}

	page := bdd.Page{Categories: categories, Posts: posts}

	// Appel de fonction qui créera la page post
	files := findPathFiles("./templates/posts.html")

	ts, err := template.ParseFiles(files...)

	// Gestion erreur 500
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, page)

}

func PostsContent(w http.ResponseWriter, r *http.Request) {
	// Gestion erreur 404
	if r.URL.Path != "/posts/content" {
		CodeErreur(w, r, 404)
		return
	}

	// Récupération de certains éléments de l'URL
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

	// Vérification de la connexion de l'utilisateur
	if VerifyCookie(w, r) {
		username := bdd.GetProfil(w, r).Username

		for i, commentEl := range comments {
			commentEl.LikeInfo = bdd.IsLiked("comment", username, commentEl.Id_comment)

			comments[i] = commentEl
		}
	}

	page := bdd.Page{Posts: posts, Comments: comments}

	files := findPathFiles("./templates/post_content.html")

	// Appel de la fonction qui créera la page post_content
	ts, err := template.ParseFiles(files...)

	// Gestion erreur 500
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	ts.Execute(w, page)

}
