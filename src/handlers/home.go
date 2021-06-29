package handlers

import (
	"fmt"
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

	filtre := r.PostFormValue("filtre")

	fmt.Println(filtre)
	// AddUser("JohnBibi", "Coucou21", "john.bibi@yforum.com")
	// DeleteUser("JohnBibi")
	// DataBase()

	// SortDate()

	// fmt.Println(filtreLike)

	// FiltresCategory()

	if len(filtre) <= 0 {
		page := bdd.Page{Posts: item, Categories: bdd.GetCategory(20, 0)}
		ts.Execute(w, page)
	} else {
		filtres := FiltreHome(w, r, filtre)

		ts.Execute(w, filtres)

	}
}

func FiltreHome(w http.ResponseWriter, r *http.Request, filtre string) bdd.Page {
	var page bdd.Page

	if filtre == "likecroissant" {
		filtres := FiltresLikeCroissant()

		page = bdd.Page{Posts: filtres, Categories: bdd.GetCategory(20, 0)}

	}

	if filtre == "likedecroissant" {
		filtres := FiltresLikeDecroissant()

		page = bdd.Page{Posts: filtres, Categories: bdd.GetCategory(20, 0)}

	}

	if filtre == "datefiltre" {
		filtres := SortDate()

		page = bdd.Page{Posts: filtres, Categories: bdd.GetCategory(20, 0)}

	}

	fmt.Println(page.Posts)

	return page

	// fmt.Println(page.Categories)

}
