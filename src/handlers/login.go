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

		passwordDB := GetPassWord(name)

		errHashed := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordForm))

		//DataBase()

		if errHashed != nil {
			fmt.Println(errHashed)
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
		} else {
			fmt.Println("right PW ", passwordDB)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else if r.URL.Path == "/login/inscription" {

		name := r.PostFormValue("registerName")
		password := r.PostFormValue("registerPassword")
		email := r.PostFormValue("registerMail")

		hashedPW := PasswordHash(password)

		AddUser(name, hashedPW, email)
		//DataBase()

	} else {
		CodeErreur(w, r, 404)
	}

	// if r.URL.Path != "/login/" {
	// 	CodeErreur(w, r, 404)
	// }

	// files := findPathFiles("./templates/login.html")

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	CodeErreur(w, r, 500)
	// }

	// ts.Execute(w, nil)
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := r.PostFormValue("registerName")
	passWord := r.PostFormValue("registerPassword")
	passWordConfirmation := r.PostFormValue("registerConfirmPassword")
	verificationDoublePassword(name, passWord, passWordConfirmation)
	// Redirection page d'accueil
	http.Redirect(w, r, "/login/", http.StatusSeeOther)
}

func PasswordAccess(name, passWord string) {
	if VerificationNumberElementPassword(passWord) {
		passWordsecure := PasswordHash(passWord)
		doublePassWordSercure := PasswordHash(passWordsecure)
		var result string = "\n Votre login est " + name + " et votre mot de passe est " + passWord
		var resultHash string = "\n Votre login est " + name + " et votre mot de passes hashé est " + passWordsecure
		var doubleResultHash string = "\n Votre login est " + name + " et votre mot de passe hashé est " + doublePassWordSercure
		fmt.Println(string("\033[1;37m\033[0m"), result)
		fmt.Println(string("\033[1;37m\033[0m"), resultHash)
		fmt.Println(string("\033[1;37m\033[0m"), doubleResultHash)

	} else {
		var redColor = "\033[31m"
		fmt.Println(string(redColor), "Votre mot de passe n'a pas passé tout les tests. \n \n")
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

func VerificationNumberElementPassword(password string) bool {
	var boolresult bool = true
	redColor := "\033[31m"
	greenColor := "\033[32m"
	// searchValue := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	searchMaj := `(.*[a-z]){2,}`
	searchMin := `(.*[A-Z]){2,}`
	searchDigit := `(.*\d){2,}`
	searchSpeChar := `(.*[@$!%*?&])`

	// re := regexp.MustCompile(searchValue)

	regexMaj := regexp.MustCompile(searchMaj)
	regexMin := regexp.MustCompile(searchMin)
	regexDigit := regexp.MustCompile(searchDigit)
	regexSpeChar := regexp.MustCompile(searchSpeChar)

	if regexMaj.Match([]byte(password)) && regexMin.Match([]byte(password)) && regexDigit.Match([]byte(password)) && regexSpeChar.Match([]byte(password)) {
		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de ponctuation.")
	} else {
		fmt.Println(string(redColor), "❌  Error : Votre mot de passe ne contient pas assez d'élément")
		boolresult = false
	}
	return boolresult
}

func verificationDoublePassword(name, firstpwd, secondpwd string) {
	if firstpwd == secondpwd {
		PasswordAccess(name, firstpwd)
	} else {
		fmt.Println("Vos deux mot de passe ne sont pas identiques \n \n")
	}
}
