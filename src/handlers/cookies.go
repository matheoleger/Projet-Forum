package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

func createCookie() {

}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "heyy",
		Value: "how many money ?",
		Path:  "/",
	})
	println("\033[0;32m", "[cookies] : we cooked your cookies, yummy !")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("heyy")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		println("\033[1;31m", "[cookies] : reading error", err)
		return
	}
	println("\033[0;32m", "[cookies] : here are the chocolat chips in your cookies :", c.Value)
}

func expire(w http.ResponseWriter, r *http.Request) {
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

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	var stringID string
	if err != nil {
		id, err := uuid.NewRandom()
		if err != nil {
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
