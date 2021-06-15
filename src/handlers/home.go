package handlers

import (
	"net/http"
	"text/template"
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

	// AddUser("JohnBibi", "Coucou21", "john.bibi@yforum.com")
	// DeleteUser("JohnBibi")
	// DataBase()

	// SessionCookie(w, r)
	// CreateCookie(w, r, "salu", "how many money")

	LaunchSession(w, r, "Alex77")

	ts.Execute(w, nil)
}
