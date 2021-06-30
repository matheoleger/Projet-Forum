package handlers

import (
	"log"
	"net/http"
	"strconv"

	"time"

	bdd "../database"
)

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

		user := bdd.GetProfil(w, r)

		InsertPost(title, content, user.Username, time.Now(), 0)

		id := GetLastedID()

		InsertBridge(id, category)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	content := r.PostFormValue("postcontent")
	postint := r.URL.Query().Get("id_post")

	n, _ := strconv.Atoi(postint)

	if !VerifyCookie(w, r) {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	} else {
		user := bdd.GetProfil(w, r)

		insertComment(content, user.Username, n)
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}

}
