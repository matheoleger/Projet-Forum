package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/login/" {

		files := findPathFiles("./templates/login.html")

		ts, err := template.ParseFiles(files...)
		if err != nil {
			CodeErreur(w, r, 500)
		}

		ts.Execute(w, nil)

	} else if r.URL.Path == "/login/connexion" {

		name := r.PostFormValue("loginName")
		passwordForm := r.PostFormValue("loginPassword")

		nameDB := GetElement(name, "username")
		passwordDB := GetElement(name, "password")

		if nameDB == name {
			errHashed := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordForm))

			// DataBase()

			if errHashed != nil {
				fmt.Println(errHashed)
				http.Redirect(w, r, "/login/?loginForm=inscription&err=wrong_PW", http.StatusSeeOther)
			} else {
				fmt.Println("right PW : ", passwordDB)
				// ExpireSession(w, r)
				LaunchSession(w, r, name)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		} else {
			http.Redirect(w, r, "/login/?loginForm=inscription&err=wrong_name", http.StatusSeeOther)
		}

	} else if r.URL.Path == "/login/inscription" {

		name := r.PostFormValue("registerName")
		password := r.PostFormValue("registerPassword")
		secondPassword := r.PostFormValue("registerConfirmPassword")
		email := r.PostFormValue("registerMail")

		if VerificationPassword(password) && VerificationEmail(email) {
			if password == secondPassword {
				hashedPW := PasswordHash(password)

				AddUser(name, hashedPW, email)
				http.Redirect(w, r, "/login/", http.StatusSeeOther)
			} else {
				fmt.Println("Votre mot de passe de confirmation n'est pas le même que votre mot de passe")
				http.Redirect(w, r, "/login/inscription", http.StatusSeeOther)
			}

		} else {
			fmt.Println("Mot de passe n'est pas bon")
			http.Redirect(w, r, "/login/?loginForm=inscription&err=wrong_PW", http.StatusSeeOther)
		}

		http.Redirect(w, r, "/login/?loginForm=connexion", http.StatusSeeOther)

	} else {
		CodeErreur(w, r, 404)
	}

}

func PasswordHash(password string) string {
	//Trasnformation du mot de passe en tableau de byte
	passWordByte := []byte(password)

	//On hash le mot de passe
	hash, err := bcrypt.GenerateFromPassword(passWordByte, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func VerificationPassword(password string) bool {

	result := true
	searchMaj := `(.*[a-z]){2,}`
	searchMin := `(.*[A-Z]){2,}`
	searchDigit := `(.*\d){2,}`
	searchSpeChar := `(.*[@$!%*?&])`
	searchLen := `[A-Za-z\d@$!%*?&]{8,}`

	regexMaj := regexp.MustCompile(searchMaj)
	regexMin := regexp.MustCompile(searchMin)
	regexDigit := regexp.MustCompile(searchDigit)
	regexSpeChar := regexp.MustCompile(searchSpeChar)
	regexLength := regexp.MustCompile(searchLen)

	if !regexMaj.Match([]byte(password)) || !regexMin.Match([]byte(password)) {
		fmt.Println("Le nombre de Majuscule et minuscule doivent être supérieurs à 2")
		result = false
	}
	if !regexDigit.Match([]byte(password)) || !regexSpeChar.Match([]byte(password)) {
		fmt.Println("Le nombre de chiffre doit être supérieurs à 2 et il doit y avoir au moins 1 caractères spéciales")
		result = false
	}
	if !regexLength.Match([]byte(password)) {
		fmt.Println("La longueur doit être de 8 caractères")
		result = false
	}

	return result
}

func VerificationEmail(email string) bool {
	var re = regexp.MustCompile(`(?mi)[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}`)

	if re.Match([]byte(email)) {
		return true
	} else {
		fmt.Println("Email non valide")
		return false
	}
}
