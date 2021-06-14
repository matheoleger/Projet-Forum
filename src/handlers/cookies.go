package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

// func CreateCookie() {

// }

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "heyy",
		Value: "how many money ?",
		Path:  "/",
	})
	println("\033[0;32m", "[cookies] : we cooked your cookies, yummy !")
}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("heyy")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		println("\033[1;31m", "[cookies] : reading error", err)
		return
	}
	println("\033[0;32m", "[cookies] : here are the chocolat chips in your cookies :", c.Value)
}

func ExpireCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("heyy")
	if err != nil {
		http.Redirect(w, r, "/nop", http.StatusSeeOther) // /set
		println("\033[1;31m", "[cookies] : expire error :", err)
		return
	}
	println("\033[0;32m", "[cookies] : all the cookies where ate :")
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func SessionCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	var stringID string
	if err != nil {
		println("\033[0;96m", "[cookies] : can't find session cookies :", err)
		id, err := uuid.NewRandom()
		if err != nil {
			println("\033[1;31m", "[cookies] : can't create uuid :", err)
		}
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			Secure:   true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)

		stringID = id.String()
	}
	println("\033[0;32m", "[cookies] : we created your session cookies : ", stringID)
}

func ExpireSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		println("\033[1;31m", "[cookies] : expire session error :", err)
		return
	}
	println("\033[0;32m", "[cookies] : all the session cookies where ate :")
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
}
