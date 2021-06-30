package handlers

import (
	"log"
	"net/http"
	"strconv"

	"time"

	bdd "../database"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	// Vérification d'une connexion de l'utilisateur
	if !VerifyCookie(w, r) {
		// Redirection vers la page de connexion si l'utilisateur n'est pas connecté
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		// Récupération des information html envoyé par l'utilisateur
		title := r.PostFormValue("titlepost")
		category := r.PostFormValue("category")
		content := r.PostFormValue("postcontent")

		user := bdd.GetProfil(w, r)

		// Intégration du post dans la base de donnée
		InsertPost(title, content, user.Username, time.Now(), 0)

		id := GetLastedID()

		// Intégration de la relation entre le post et sa catégorie dans la base de donnée
		InsertBridge(id, category)

		// Redirection vers la page d'accueil
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Récupération des information html envoyé par l'utilisateur ainsi que d'un élément de l'URL
	content := r.PostFormValue("postcontent")
	postint := r.URL.Query().Get("id_post")

	n, _ := strconv.Atoi(postint)

	// Vérification d'une connexion de l'utilisateur
	if !VerifyCookie(w, r) {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	} else {
		user := bdd.GetProfil(w, r)

		// Insertion du commentaire
		insertComment(content, user.Username, n)

		// Redirection vers la page précédemment visité
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}

}
