package handlers

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//fonction qui créé un cookie de session avec SessionCookie()
//puis ajoute une session à la base de donnée avec AddSession()
func LaunchSession(w http.ResponseWriter, r *http.Request, username string) {
	uuid := SessionCookie(w, r)
	if uuid == "error" {
		println("\033[1;31m", "[session] : launching session error : session, probably allready exist")
		return

	} else {
		AddSession(uuid, username)
	}
}

// comme LaunchSession() mais supprime le cookie et la session dans la database
func EndSession(w http.ResponseWriter, r *http.Request) {
	uuid := ExpireSession(w, r)
	if uuid == "error" {
		println("\033[1;31m", "[session] : expire session error ")
		return

	} else {
		DeleteSession(uuid)
	}
}

//fonction qui ajoute une session dans la database, qui correspond à un lien entre un uuid et un utilisateur
func AddSession(uuid string, user_name string) {
	db := OpenDataBase()

	//cherche un username dans la base de donnée qui à pour uuid celui que l'on s'aprête a créer
	statementuuid, err4 := db.Prepare("SELECT username FROM session WHERE uuid = ?")
	if err4 != nil {
		fmt.Println("\033[1;31m", "[session] : error, impossible to do statement")
		return
	}

	resultuuid, err5 := statementuuid.Query(uuid)
	if err5 != nil {
		fmt.Println("\033[1;31m", "[session] : error, impossible to query")
		return
	}

	var resultname string
	for resultuuid.Next() { //essaie de récuperer les données de cette requette
		resultuuid.Scan(&resultname)
		resultuuid.Close()
		fmt.Println("\033[1;31m", "[session] : error, a session allready exist with uuid =", uuid)
		//si il trouve des valeurs, cela veut dire qu'une session existe déjà avec cet uuid,
		// donc la session que l'on essaie de creer est invalide, donc on return vide

		return
	}

	//même chose mais cette fois-ci on cherche un uuid en fonction d'e l'username
	statementTest, err := db.Prepare("SELECT uuid FROM session WHERE username = ?")
	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, impossible to verify if session allready exist", uuid)
		return
	}
	result, err2 := statementTest.Query(user_name)

	if err2 != nil {
		fmt.Println("\033[1;31m", "[session] : error, impossible to do statement :", uuid, "\n", result)
		return
	}

	var resultId string
	for result.Next() {
		result.Scan(&resultId) //si on trouve des valeurs, ça veut dire qu'une session existe déjà avec cet username
		result.Close()
		println("\033[0;96m", "[session] : there is allready a session with this username :")
		//mais plutot que d'abandonner, on met a jour la session

		statementUpdate, err3 := db.Prepare("UPDATE session SET uuid = ? WHERE username = ?")
		if err3 != nil {
			fmt.Println("\033[1;31m", "[session] : error, impossible update session :", uuid)
			return
		}
		statementUpdate.Exec(uuid, user_name) //donc on update avec le nouvel uuid

		return

	}

	//et si aucun des deux n'a de données, cela veut dire que l'on peut créer la session
	statement, err := db.Prepare("INSERT INTO session (uuid, username) VALUES (?, ?)")

	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, can't insert into database")
		return
	}

	statement.Exec(uuid, user_name)

}

//fonction qui supprime une session en fonction d'un uuid
func DeleteSession(uuid string) {
	db := OpenDataBase()

	statement, err := db.Prepare("DELETE FROM session WHERE uuid = ?")

	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, deleting from database")
		return
	}
	statement.Exec(uuid)
}
