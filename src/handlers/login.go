package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"

	bdd "../database"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Gestion erreur 404
	if r.URL.Path == "/login/" {

		// Appel de la fonction qui créera la page login
		files := findPathFiles("./templates/login.html")

		ts, err := template.ParseFiles(files...)

		// Gestion d'erreur 500
		if err != nil {
			CodeErreur(w, r, 500)
		}

		ts.Execute(w, nil)

	} else if r.URL.Path == "/login/connexion" {

		// Récupération des information entrées par l'utilisateur
		name := r.PostFormValue("loginName")
		passwordForm := r.PostFormValue("loginPassword")

		// Récupération des informations de la base de données
		nameDB := GetElement(name, "username")
		passwordDB := GetElement(name, "password")

		// Vérification du nom entré par l'utilisateur avec les noms de la base de données
		if nameDB == name {

			// Comparaison des mots de passes hashés de la base de données avec l'utilisateur
			errHashed := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordForm))

			if errHashed != nil {
				http.Redirect(w, r, "/login/?loginForm=inscription&err=wrong_PW", http.StatusSeeOther)
			} else {
				// Création de session si toutes les conditions sont validés
				LaunchSession(w, r, name)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		} else {
			http.Redirect(w, r, "/login/?loginForm=inscription&err=wrong_name", http.StatusSeeOther)
		}

	} else if r.URL.Path == "/login/inscription" {
		// Récupération de chaques informations entrées par l'utilisateur
		name := r.PostFormValue("registerName")
		password := r.PostFormValue("registerPassword")
		secondPassword := r.PostFormValue("registerConfirmPassword")
		email := r.PostFormValue("registerMail")

		// Verification de chaque critères pour la création d'un compte
		if VerificationPassword(password) && VerificationEmail(email) {
			test := bdd.GetInformationAllUser(w, r, "email")
			if bdd.VerificationEmail(email, test) {
				if password == secondPassword {
					hashedPW := PasswordHash(password)
					// Ajout d'un utilisateur si et seulement si tout les critères sont valides
					AddUser(name, hashedPW, email)
					http.Redirect(w, r, "/login/", http.StatusSeeOther)
				} else {
					fmt.Println("Votre mot de passe de confirmation n'est pas le même que votre mot de passe")
					http.Redirect(w, r, "/login/inscription", http.StatusSeeOther)
				}
			} else {
				fmt.Println("Un utilisateur déjà inscrit possède votre addresse mail")
				http.Redirect(w, r, "/login/?loginForm=inscription&err=wrong_EMAIL", http.StatusSeeOther)
			}

		} else {
			fmt.Println("Mot de passe n'est pas bon")
			http.Redirect(w, r, "/login/?loginForm=inscription&err=wrong_PW", http.StatusSeeOther)
		}
		// Si une information est invalide alors l'utilisateur restera sur la même page
		http.Redirect(w, r, "/login/?loginForm=connexion", http.StatusSeeOther)

	} else {
		// Gestion d'erreur 404
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

	// Utilisation de regex afin de vérifier chaque élément du mot de passe
	result := true
	searchMaj := `(.*[a-z]){2,}`
	searchMin := `(.*[A-Z]){2,}`
	searchDigit := `(.*\d){2,}`
	searchSpeChar := `(.*\W|_)`
	searchLen := `[A-Za-z\d@$!%*?&]{8,}`

	regexMaj := regexp.MustCompile(searchMaj)
	regexMin := regexp.MustCompile(searchMin)
	regexDigit := regexp.MustCompile(searchDigit)
	regexSpeChar := regexp.MustCompile(searchSpeChar)
	regexLength := regexp.MustCompile(searchLen)

	// Message d'erreur si mot de passe non valide
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

	// Vérification de chaque charactère du mail
	var re = regexp.MustCompile(`(?mi)[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}`)

	if re.Match([]byte(email)) {
		return true
	} else {
		fmt.Println("Email non valide")
		return false
	}
}
