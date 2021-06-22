package handlers

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

func CreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.PostFormValue("titlepost")
	category := r.PostFormValue("category")
	content := r.PostFormValue("postcontent")
	postFile := r.PostFormValue("postfile")

	content += postFile

	f, err := os.Open(postFile)
	if err != nil {
		fmt.Println("error open file")
	}

	image, _, err := image.Decode(f)

	fmt.Println(image)

	fmt.Println("Votre titre est : " + title + " et votre catégorie est : " + category + " puis votre contenu est : " + content)

	username := "Johanna"

	InsertPost(title, content, username, 0, false, time.Now())

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	content := r.PostFormValue("postcontent")
	post := r.PostFormValue("postfile")
	n, _ := strconv.Atoi(post)

	fmt.Println(n)
	fmt.Println("Votre contenu est : " + content + "sur le post : " + post)

	username := "Johanna"

	insertComment(content, username, n)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
