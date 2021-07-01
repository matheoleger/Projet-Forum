package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

//Fonction qui retourne vrai si un cookie de session existe et faux si il n'existe pas
func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("session")
	if err != nil {
		return false
	} else {
		return true
	}
}

//Fonction qui vérifie si un cookie de session existe, si il n'existe pas alors, il va en créer un
func SessionCookie(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("session") //try to read cookie

	//Vérification de la présence d'un cookie de session
	if !VerifyCookie(w, r) {
		println("\033[0;96m", "[cookies] : can't find session cookies :", err)
		id, err2 := uuid.NewRandom() // Création nouveau uuid
		if err2 != nil {
			println("\033[1;31m", "[cookies] : can't create uuid :", err)
		}
		// Création d'un nouveau cookie
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			Secure:   true,
			HttpOnly: true,
			Expires:  time.Now().Add(2 * time.Hour),
			Path:     "/",
		}
		http.SetCookie(w, cookie)
		return *&cookie.Value // Création du cookie + renvoie de sa valeur

	}
	println("\033[1;31m", "[cookies] : session allready exist")
	return "error"
}

// Suppression de la session lors qu'elle expire
func ExpireSession(w http.ResponseWriter, r *http.Request) string {
	c, err := r.Cookie("session") // Récupération du cookie
	if err != nil {
		println("\033[1;31m", "[cookies] : expire session error :", err)
		return "error"
	}

	c.MaxAge = -1 // Suppression du cookie
	http.SetCookie(w, c)
	return *&c.Value
}
