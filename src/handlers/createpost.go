package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Creationpost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/creationpost" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/createpost.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}
	ts.Execute(w, nil)

}

func GetElementOfPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.PostFormValue("titlepost")
	category := r.PostFormValue("category")
	content := r.PostFormValue("postcontent")

	fmt.Println("Votre titre est :" + title + " vous catégorie est " + category + " votre contenu est " + content)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
