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
				http.Redirect(w, r, "/login/", http.StatusSeeOther)
			} else {
				fmt.Println("right PW ", passwordDB)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}

	} else if r.URL.Path == "/login/inscription" {

		name := r.PostFormValue("registerName")
		password := r.PostFormValue("registerPassword")
		email := r.PostFormValue("registerMail")

		if VerificationPassword(password) && VerificationEmail(email) {
			hashedPW := PasswordHash(password)

			AddUser(name, hashedPW, email)
			// DataBase()
		} else {
			fmt.Println("Mot de passe n'est pas bon")
		}

		http.Redirect(w, r, "/login/", http.StatusSeeOther)

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

// func GetLogin(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	name := r.PostFormValue("loginName")
// 	passWord := r.PostFormValue("loginPassword")
// 	PasswordAccess(name, passWord)

// 	// Redirection page d'accueil
// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
// }

// func PasswordAccess(name, passWord string) {
// 	if SecurisationPassword(passWord) {
// 		passWordsecure := PasswordHash(passWord)
// 		doublePassWordSercure := PasswordHash(passWordsecure)
// 		var result string = "\n Votre login est " + name + " et votre mot de passe est " + passWord
// 		var resultHash string = "\n Votre login est " + name + " et votre mot de passes hashé est " + passWordsecure
// 		var doubleResultHash string = "\n Votre login est " + name + " et votre mot de passe hashé est " + doublePassWordSercure
// 		fmt.Println(string("\033[1;37m\033[0m"), result)
// 		fmt.Println(string("\033[1;37m\033[0m"), resultHash)
// 		fmt.Println(string("\033[1;37m\033[0m"), doubleResultHash)
// 	} else {
// 		var redColor = "\033[31m"
// 		fmt.Println(string(redColor), "Votre mot de passe n'a pas passé tout les tests. \n \n")
// 	}
// }

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

//On compte le nombre d'élément obligatoire
// func SecurisationPassword(passWord string) bool {
// 	passWordRune := []rune(passWord)
// 	punct := 0
// 	num := 0
// 	maj := 0
// 	min := 0

// 	if len(passWord) <= 8 {
// 		fmt.Println("⚠️  Votre mot de passe n'est pas assez long, votre mot de passe contient " + strconv.Itoa(len(passWord)) + " caractères , veuillez réessayer")
// 	} else {
// 		for indexOfPassWord := 0; indexOfPassWord < len(passWord); indexOfPassWord++ {
// 			//Verification nombre de ponctuation
// 			for indexOfRune := 0; indexOfRune < len(passWordRune); indexOfRune++ {
// 				if unicode.IsPunct(passWordRune[indexOfRune]) {
// 					punct++
// 				}
// 			}
// 			//Vérification nombre de chiffre
// 			if passWord[indexOfPassWord] >= 48 && passWord[indexOfPassWord] <= 57 {
// 				num++
// 			}
// 			//Vérification nombre de majuscule
// 			if passWord[indexOfPassWord] >= 65 && passWord[indexOfPassWord] <= 90 {
// 				maj++
// 			}
// 			//Vérification nombre de minuscule
// 			if passWord[indexOfPassWord] >= 97 && passWord[indexOfPassWord] <= 122 {
// 				min++
// 			}
// 		}
// 	}
// 	return VerificationNumberElementPassword(punct, num, maj, min)
// }

// func VerificationNumberElementPassword(punct, num, maj, min int) bool {
// 	var result bool = true
// 	var redColor = "\033[31m"
// 	var greenColor = "\033[32m"
// 	if punct < 1 {
// 		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de ponctuation.")
// 		result = false
// 	} else {
// 		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de ponctuation.")
// 	}
// 	if num < 1 {
// 		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de chiffre.")
// 		result = false
// 	} else {
// 		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de chiffres.")
// 	}
// 	if maj < 2 {
// 		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de majuscule.")
// 		result = false
// 	} else {
// 		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de majuscule.")
// 	}
// 	if min < 2 {
// 		fmt.Println(string(redColor), "⚠️  Error : Votre mot de passe ne contient pas assez de minuscule.")
// 		result = false
// 	} else {
// 		fmt.Println(string(greenColor), "✔️  Votre mot de passe contient assez de minuscule.")
// 	}
// 	return result
// }

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

	// testValue := "HeHe12?"

	// if regexMaj.Match([]byte(testValue)) && regexMin.Match([]byte(testValue)) && regexDigit.Match([]byte(testValue)) && regexSpeChar.Match([]byte(testValue)) {
	// 	fmt.Println("everything is fine ;)")
	// } else {
	// 	fmt.Println("wrong duuuude")
	// }

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
