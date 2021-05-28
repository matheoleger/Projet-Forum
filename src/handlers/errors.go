package handlers

import (
	"fmt"
	"net/http"
)

func CodeErreur(w http.ResponseWriter, r *http.Request, status int, message string) {
	const colorRed = "\033[31m"

	if status == 404 {
		http.Error(w, "404 not found", http.StatusNotFound)
		fmt.Println(message)
	}
	if status == 400 {
		http.Error(w, "400 Bad request", http.StatusBadRequest)
		fmt.Println(message)
	}
	if status == 500 {
		http.Error(w, "500 Internal Server", http.StatusInternalServerError)
		fmt.Println(message)
	}
}
