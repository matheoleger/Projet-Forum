package handlers

import (
	"net/http"
	"text/template"

	bdd "../database"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Gestion d'erreur 404
	if r.URL.Path != "/" {
		CodeErreur(w, r, 404)
		return
	}

	// Appel de la fonction qui créera la page d'accueil
	files := findPathFiles("./templates/home.html")

	ts, err := template.ParseFiles(files...)

	// Gestion d'erreur 500
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	// Item récupère toutes les informations des posts
	item := GetPost()

	if VerifyCookie(w, r) {
		username := bdd.GetProfil(w, r).Username

		for i, postEl := range item {
			postEl.LikeInfo = bdd.IsLiked("post", username, postEl.Id_post)

			item[i] = postEl
		}
	}

	// Récupération du choix du filtre
	filtre := r.PostFormValue("filtre")

	if len(filtre) <= 0 {

		// Affichage des posts par défault, sans aucun filtre
		page := bdd.Page{Posts: item, Categories: bdd.GetCategory(20, 0)}
		ts.Execute(w, page)
	} else {

		// Affichage des posts avec un filtre
		filtres := FiltreHome(w, r, filtre)
		ts.Execute(w, filtres)

	}
}

func FiltreHome(w http.ResponseWriter, r *http.Request, filtre string) bdd.Page {
	var page bdd.Page

	var filtres []bdd.Post

	// Appel de fonction permettant le filtre en fonction du choix de l'utilisateur
	if filtre == "likecroissant" {
		filtres = FiltresLikeDecroissant()

	} else if filtre == "likedecroissant" {
		filtres = FiltresLikeCroissant()

	} else if filtre == "datefiltre" {
		filtres = SortDate()

	} else {
		CodeErreur(w, r, 400)
	}

	if VerifyCookie(w, r) {
		username := bdd.GetProfil(w, r).Username

		for i, postEl := range filtres {
			postEl.LikeInfo = bdd.IsLiked("post", username, postEl.Id_post)

			filtres[i] = postEl
		}
	}

	page = bdd.Page{Posts: filtres, Categories: bdd.GetCategory(20, 0)}

	return page

}
