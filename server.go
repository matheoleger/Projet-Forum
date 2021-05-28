package main

import (
	"net/http"

	handlers "./src/handlers"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) // récupère tous les fichiers "externe" dans "static"

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/categories", handlers.Categories)
	http.HandleFunc("/posts/", handlers.Posts)
	http.HandleFunc("/profil/", handlers.Profil)

	http.ListenAndServe(":8080", nil)
}
