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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Génère des pages grâce aux fonctions du package Handlers
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/categories", handlers.Categories)
	http.HandleFunc("/posts", handlers.Posts)
	http.HandleFunc("/posts/content", handlers.PostsContent)
	http.HandleFunc("/profil", handlers.Profil)
	http.HandleFunc("/login/", handlers.Login)
	http.HandleFunc("/like", handlers.Like)

	http.HandleFunc("/creationpost", handlers.CreatePost)

	fmt.Println(string(white), "[SERVER_READY] : on http://localhost:8080 ✅ ")
	fmt.Println(string(yellow), "[SERVER_INFO] : To stop the program : Ctrl + c", string(white))

	// Lecture du serveur sur le port 8080
	http.ListenAndServe(":8080", nil)
}
