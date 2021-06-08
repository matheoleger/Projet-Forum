package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/login" {

		files := findPathFiles("./templates/login.html")

		ts, err := template.ParseFiles(files...)
		if err != nil {
			CodeErreur(w, r, 500)
		}

		ts.Execute(w, nil)

		// } else if r.URL.Path == "/login/connexion" {

		// 	name := r.PostFormValue("loginName")
		// 	passwordForm := r.PostFormValue("loginPassword")

		// 	passwordDB := GetPassWord(name)

		// 	errHashed := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordForm))

		// 	DataBase()

		// 	if errHashed != nil {
		// 		fmt.Println(errHashed)
		// 		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		// 	} else {
		// 		fmt.Println("right PW ", passwordDB)
		// 		http.Redirect(w, r, "/", http.StatusSeeOther)
		// 	}

		// } else if r.URL.Path == "/login/inscription" {

		// 	name := r.PostFormValue("registerName")
		// 	password := r.PostFormValue("registerPassword")
		// 	email := r.PostFormValue("registerMail")

		// 	hashedPW := PasswordHash(password)

		// 	AddUser(name, hashedPW, email)
		// 	DataBase()

		// }
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
	name := r.PostFormValue("loginName")
	passWord := r.PostFormValue("loginPassword")
	SecurisationPassword(passWord)
	passWordsecure := PasswordHash(passWord)
	doublePassWordSercure := PasswordHash(passWordsecure)
	var result string = "\n Votre login est " + name + " et votre mot de passe est " + passWord
	var resultHash string = "\n Votre login est " + name + " et votre mot de passes hashé est " + passWordsecure
	var doubleResultHash string = "\n Votre login est " + name + " et votre mot de passe hashé est " + doublePassWordSercure

	if len(name) == 0 && len(passWord) == 0 {
		fmt.Println("⚠️ Votre mot de passe n'a pas été enregistré")
	} else {
		fmt.Println(string("\033[1;37m\033[0m"), result)
		fmt.Println(string("\033[1;37m\033[0m"), resultHash)
		fmt.Println(string("\033[1;37m\033[0m"), doubleResultHash)
	}

	// Redirection page d'accueil
	http.Redirect(w, r, "/login", http.StatusSeeOther)
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

func SecurisationPassword(passWord string) {
	passWordRune := []rune(passWord)
	punct := 0
	maj := 0
	min := 0
	num := 0

	if len(passWord) <= 8 {
		fmt.Println("⚠️  Votre mot de passe n'est pas assez long, votre mot de passe contient " + strconv.Itoa(len(passWord)) + " caractères , veuillez réessayer")
	} else {
		for indexOfPassWord := 0; indexOfPassWord < len(passWord); indexOfPassWord++ {
			//Verification nombre de ponctuation
			for indexOfRune := 0; indexOfRune < len(passWordRune); indexOfRune++ {
				if unicode.IsPunct(passWordRune[indexOfRune]) {
					punct++
				}
			}
			//Vérification nombre de chiffre
			if passWord[indexOfPassWord] >= 48 && passWord[indexOfPassWord] <= 57 {
				num++
			}
			//Vérification nombre de majuscule
			if passWord[indexOfPassWord] >= 65 && passWord[indexOfPassWord] <= 90 {
				maj++
			}
			//Vérification nombre de minuscule
			if passWord[indexOfPassWord] >= 97 && passWord[indexOfPassWord] <= 122 {
				min++
			}
		}
	}
	VerificationPassWord(punct, num, maj, min)
}

func VerificationPassWord(punct, num, maj, min int) {
	var redColor = "\033[31m"
	var greenColor = "\033[32m"
	if punct < 1 {
		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de ponctuation.")
	} else {
		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de ponctuation.")
	}
	if num <= 2 {
		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de chiffre.")
	} else {
		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de chiffres.")
	}
	if maj <= 1 {
		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de majuscule.")
	} else {
		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de majuscule.")
	}
	if min <= 1 {
		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de minuscule.")
	} else {
		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de minuscule.")
	}
}
