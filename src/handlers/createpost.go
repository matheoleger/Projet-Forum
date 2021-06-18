package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// func Creationpost(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/creationpost" {
// 		CodeErreur(w, r, 404)
// 		return
// 	}

// 	files := findPathFiles("./templates/createpost.html")

// 	ts, err := template.ParseFiles(files...)
// 	if err != nil {
// 		CodeErreur(w, r, 500)
// 		return
// 	}
// 	ts.Execute(w, nil)

// }

// func CreatePost(w http.ResponseWriter, r *http.Request) {

// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	title := r.PostFormValue("titlepost")
// 	category := r.PostFormValue("category")
// 	content := r.PostFormValue("postcontent")

// 	fmt.Println("Votre titre est : " + title + " et votre catégorie est : " + category + " puis votre contenu est : " + content)

// 	username := "Johanna"

// 	InsertPost(title, content, username)

// 	http.Redirect(w, r, "/", http.StatusSeeOther)

// }

func createComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	post := r.PostFormValue("commentpost")
	content := r.PostFormValue("commentcontent")

	fmt.Println("Votre contenu est : " + content + "sur le post : " + post)

	username := "Johanna"

	insertComment(content, username, post)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
