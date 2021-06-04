package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		CodeErreur(w, r, 404, "[Server_Alert] - Error 404 : Page not found")
	}

	files := findPathFiles("./templates/home.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500, "[Server_Alert] - Error 500 : Template not found -> home.html")
	}

	AddUser("JohnBibi", "Coucou21", "john.bibi@yforum.com")
	DataBase()
	
	ts.Execute(w, nil)
}
