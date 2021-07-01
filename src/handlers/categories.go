package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	bdd "../database"
)

func Categories(w http.ResponseWriter, r *http.Request) {
	// Gestion d'erreur 404
	if r.URL.Path != "/categories" {
		CodeErreur(w, r, 404)
		return
	}

	// Appel de la fonction qui créera la page d'accueil
	files := findPathFiles("./templates/categories.html")

	ts, err := template.ParseFiles(files...)

	// Gestion d'erreur 500
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	// Récupération d'élément de l'URL
	wichpage := r.URL.Query().Get("page")
	wichpageInt, _ := strconv.Atoi(wichpage)

	perpage := r.URL.Query().Get("perpage")
	perpageInt, _ := strconv.Atoi(perpage)

	// Affichage des catégories
	page := bdd.Page{Categories: bdd.GetCategory(perpageInt, wichpageInt)}
	ts.Execute(w, page)
}

func InsertBridge(B_id_post int, B_name_category string) {

	// Ouverture de la base de donnée
	db := OpenDataBase()

	// Préparation d'une future association d'un post à une catégorie dans la base de donnée
	statement, err := db.Prepare("INSERT INTO bridge (B_id_post, B_name_category) VAlUES (?, ?)")

	// Gestion d'erreur de la base de donnée
	if err != nil {
		fmt.Println("error prepare ")
		return
	}

	// Exécution de la préparation
	statement.Exec(B_id_post, B_name_category)

	// Fermeture base de donnée
	defer db.Close()
}
