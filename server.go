package main

import (
	"fmt"
	"net/http"

	handlers "./src/handlers"
)

func main() {
	const white = "\033[1;37m\033[0m"
	const blue = "\033[34m"
	const yellow = "\033[33m"

	fmt.Println(string(blue), "[SERVER_INFO] : Starting local Server...")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) // récupère tous les fichiers "externe" dans "static"

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/categories", handlers.Categories)
	http.HandleFunc("/posts", handlers.Posts)
	http.HandleFunc("/profil", handlers.Profil)

	fmt.Println(string(white), "[SERVER_READY] : on http://localhost:8000 ✅ ")
	fmt.Println(string(yellow), "[SERVER_INFO] : To stop the program : Ctrl + c")
	http.ListenAndServe(":8080", nil)
}
