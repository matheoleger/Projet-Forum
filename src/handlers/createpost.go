package handlers

import (
	"fmt"
	// "image"
	"log"
	"net/http"

	// "os"
	"time"

	bdd "../database"
)

// func Creationpost(w http.ResponseWriter, r *http.Request) {
//     if r.URL.Path != "/creationpost" {
//         CodeErreur(w, r, 404)
//         return
//     }

//     files := findPathFiles("./templates/createpost.html")

//     ts, err := template.ParseFiles(files...)
//     if err != nil {
//         CodeErreur(w, r, 500)
//         return
//     }
//     ts.Execute(w, nil)

// }

func CreatePost(w http.ResponseWriter, r *http.Request) {

	if !VerifyCookie(w, r) {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		title := r.PostFormValue("titlepost")
		category := r.PostFormValue("category")
		content := r.PostFormValue("postcontent")
		// postFile := r.PostFormValue("postfile")

		user := bdd.GetProfil(w, r)

		// content += postFile

		// f, err := os.Open(postFile)
		// if err != nil {
		//     fmt.Println("error open file")
		// }

		// image, _, err := image.Decode(f)

		// fmt.Println(image)

		fmt.Println("Votre titre est : " + title + " et votre catégorie est : " + category + " puis votre contenu est : " + content + ". C'est " + user.Username + " qui a créé ce post, il y a 0 likes et ce post n'a pas été liké, de plus il a été créé le : " + time.Now().String())

		InsertPost(title, content, user.Username, time.Now(), 0)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
