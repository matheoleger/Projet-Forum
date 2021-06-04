package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		CodeErreur(w, r, 404)
	}

	files := findPathFiles("./templates/login.html")

	ts, err := template.ParseFiles(files...)
	if err != nil {
		CodeErreur(w, r, 500)
	}

	ts.Execute(w, nil)
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := r.PostFormValue("loginName")
	passWord := r.PostFormValue("loginPassword")
	passWordsecure := PasswordHash(w, r, passWord)
	doublePassWordSercure := PasswordHash(w, r, passWordsecure)
	var result string = "\n Votre login est " + name + " et votre mot de passe est " + passWord
	var resultHash string = "\n Votre login est " + name + " et votre mot de passes hashé est " + passWordsecure
	var doubleResultHash string = "\n Votre login est " + name + " et votre mot de passe hashé est " + doublePassWordSercure

	if len(name) == 0 && len(passWord) == 0 {
		fmt.Println("Votre mot de passe n'a pas été enregistré")
	} else {
		fmt.Println(string("\033[1;37m\033[0m"), result)
		fmt.Println(string("\033[1;37m\033[0m"), resultHash)
		fmt.Println(string("\033[1;37m\033[0m"), doubleResultHash)
	}
	// Redirection page d'accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func PasswordHash(w http.ResponseWriter, r *http.Request, password string) string {
	// Trasnformation du mot de passe en tableau de byte
	passWordByte := []byte(password)

	// On hash le mot de passe
	hash, err := bcrypt.GenerateFromPassword(passWordByte, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
