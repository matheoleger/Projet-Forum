package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

func CreateCookie(w http.ResponseWriter, r *http.Request, name string, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	})
	println("\033[0;32m", "[cookies] : we cooked your cookies, yummy !")
}

func ReadCookie(w http.ResponseWriter, r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		println("\033[1;31m", "[cookies] : reading error", err)
		return "empty"
	}
	println("\033[0;32m", "[cookies] : here are the chocolat chips in your cookies :", c.Value)
	return c.Value
}

func ExpireCookie(w http.ResponseWriter, r *http.Request, name string) {
	c, err := r.Cookie(name)
	if err != nil {
		//http.Redirect(w, r, "/nop", http.StatusSeeOther) // /set
		println("\033[1;31m", "[cookies] : expire error :", err)
		return
	}
	println("\033[0;32m", "[cookies] : all the cookies where ate :")
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
}

func SessionCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session") //try read cookie
	var stringID string
	//if erorr (no cookie named session)
	if err != nil {
		println("\033[0;96m", "[cookies] : can't find session cookies :", err)
		id, err := uuid.NewRandom() //create new uuid
		if err != nil {
			println("\033[1;31m", "[cookies] : can't create uuid :", err)
		}
		// create a new cookie
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			Secure:   true,
			HttpOnly: true,
			// Path:     "/",
		}
		http.SetCookie(w, cookie)

		stringID = id.String()
		println("\033[0;32m", "[cookies] : we created your session cookies : ", stringID)

		//test
		AddSession(stringID, "Johanna")
		return

	}
	println("\033[1;31m", "[cookies] : session allready exist")

}

func ExpireSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		println("\033[1;31m", "[cookies] : expire session error :", err)
		return
	}
	println("\033[0;32m", "[cookies] : all the session cookies where ate, humm yummy !")
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
}

func LaunchSession(w http.ResponseWriter, r *http.Request, username string) {
	SessionCookie(w, r)
	var uuid = ReadCookie(w, r, "session")
	println("\033[0;32m", "[session] : launch session, uuid = ", uuid)
	AddSession(uuid, username)
}

func EndSession(w http.ResponseWriter, r *http.Request) {
	ExpireSession(w, r)
	var uuid = ReadCookie(w, r, "session")
	println("\033[0;32m", "[session] : end session ")
	DeleteSession(uuid)

}
