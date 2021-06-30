package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func CodeErreur(w http.ResponseWriter, r *http.Request, status int) {

	// Récupération du contenu html de la page error
	files := findPathFiles("./templates/error.html")

	t, err := template.ParseFiles(files...)

	if err != nil {
		log.Fatal(err)
	}

	// Permet l'affichage du contenu de la page error
	t.Execute(w, ErrorType(w, r, status))

}

func ErrorType(w http.ResponseWriter, t *http.Request, status int) string {

	// Fonction renvoyant une erreur en fonction de l'erreur reconnu
	var errorstr string
	if status == 400 {
		errorstr = `Error 400 : Bad Request `
	}
	if status == 404 {
		errorstr = `Error 404 : Page Not Found`
	}
	if status == 500 {
		errorstr = `Error 500 : Internal Server Error`
	}
	return errorstr
}
