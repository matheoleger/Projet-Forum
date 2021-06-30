package handlers

import (
	"net/http"
	"text/template"

	bdd "../database"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profil" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/profil.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	Disconnect(w, r)

	//Vérification de présence de cookie de connexion
	if !VerifyCookie(w, r) {
		//Redirect vers page de connexion si non présence de cookie
		http.Redirect(w, r, "/login/", http.StatusSeeOther)

	} else {
		//Affichage de données de l'utilisateur si présence de cookie
		content := bdd.GetProfil(w, r)
		page := bdd.Page{UserInfo: content, Posts: SortUserPost(content.Username)}

		ts.Execute(w, page)
	}

}

func Disconnect(w http.ResponseWriter, r *http.Request) {
	//Récupération du bouton de déconnexion
	submit := r.FormValue("submit")

	if len(submit) != 0 {
		//Suppression du cookie de connexion et redirect vers la page login
		EndSession(w, r)
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	}

}
