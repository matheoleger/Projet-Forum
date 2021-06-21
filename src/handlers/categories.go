package handlers

import (
	"net/http"
	"text/template"
)

func Categories(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/categories" {
		CodeErreur(w, r, 404)
		return
	}

	files := findPathFiles("./templates/categories.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
		return
	}

	// ExpireSession(w, r)
	// ReadCookie(w, r, "salu")
	// ExpireCookie(w, r, "salu")
	// ExpireSession(w, r)
	LaunchSession(w, r, "Chad")
	// SessionCookie(w, r)

	// EndSession(w, r)

	ts.Execute(w, nil)
}
