package handlers

import (
	"net/http"
	"text/template"

	bdd "../database"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/home.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}
	item := GetPost()
	page := bdd.Page{Posts: item, Categories: bdd.GetCategory()}
	// AddUser("JohnBibi", "Coucou21", "john.bibi@yforum.com")
	// DeleteUser("JohnBibi")
	// DataBase()

	FiltresLike()

	// fmt.Println(filtreLike)

	FiltresCategory()

	ts.Execute(w, page)
}
