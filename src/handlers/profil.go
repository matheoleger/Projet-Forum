package handlers

import (
	"fmt"
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

	if !VerifyCookie(w, r) {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)

	} else {
		content := bdd.GetProfil(w, r)
		fmt.Println(content)

		page := bdd.Page{UserInfo: content}

		ts.Execute(w, page)
	}

}
