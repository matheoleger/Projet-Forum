package database

import (
	"fmt"
	"net/http"
)

func GetProfil(w http.ResponseWriter, r *http.Request) User {
	db := OpenDataBase()

	// Récupère le cookie de session de l'utilisateur
	c, err := r.Cookie("session")

	if err != nil {
		// Gestion erreur cookie
		fmt.Println("Cookie error")
	}
	test := c.Value

	// Séléection du nom de l'utilisateur en fonction de sa session
	statement, err := db.Prepare("SELECT username FROM session WHERE uuid = ?")

	if err != nil {
		// Message erreur BDD
		fmt.Println("Error prepare")
	}

	result, err2 := statement.Query(test)

	if err2 != nil {
		// Message erreur BDD
		fmt.Println("query error")
	}

	defer statement.Close()
	var username string

	// Parcourir les colonnes sélectionné précédement
	for result.Next() {
		result.Scan(&username)
	}

	resultInformation := GetMoreInformation(w, r, username)
	return resultInformation
}

func GetMoreInformation(w http.ResponseWriter, r *http.Request, username string) User {
	db := OpenDataBase()

	var myUser User

	// Selection du mail ainsi que du nom de l'utilisateur depuis la table user
	statement, err := db.Prepare("SELECT username, email FROM user WHERE username = ?")

	if err != nil {
		// Message d'erreur BDD
		fmt.Println("Error prepare")
	}

	result, err2 := statement.Query(username)

	if err2 != nil {
		// Message d'erreur BDD
		fmt.Println("query error")
	}

	defer statement.Close()

	for result.Next() {
		result.Scan(&myUser.Username, &myUser.Email)
	}

	return myUser
}

func GetInformationAllUser(w http.ResponseWriter, r *http.Request, element string) []string {
	db := OpenDataBase()

	var users []string

	// Selection d'une colonne spécifique depuis la table use
	statement, err := db.Query("SELECT " + element + " FROM user")

	if err != nil {
		// Message erreur BDD
		fmt.Println("Error prepares")
	}

	defer statement.Close()

	var elements string
	for statement.Next() {
		statement.Scan(&elements)

		// Ajout des élémentd dans le tableau users
		users = append(users, elements)
	}

	return users
}

func VerificationEmail(email string, emailbdd []string) bool {
	// Fonction qui va vérifier les caractères de l'email
	result := true
	for index := 0; index < len(emailbdd); index++ {
		if email == emailbdd[index] {
			result = false
			break
		}
	}
	return result
}
