package handlers

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func LaunchSession(w http.ResponseWriter, r *http.Request, username string) {
	uuid := SessionCookie(w, r)
	if uuid == "error" {
		println("\033[1;31m", "[session] : launching session error : session, probably allready exist")
		return

	} else {
		AddSession(uuid, username)
	}
}

func EndSession(w http.ResponseWriter, r *http.Request) {
	uuid := ExpireSession(w, r)
	if uuid == "error" {
		println("\033[1;31m", "[session] : expire session error ")
		return

	} else {
		DeleteSession(uuid)
	}
}

func AddSession(uuid string, user_name string) {
	db := OpenDataBase()

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
	for resultuuid.Next() {
		resultuuid.Scan(&resultname)
		resultuuid.Close()
		fmt.Println("\033[1;31m", "[session] : error, a session allready exist with uuid =", uuid)

		return
	}

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
		result.Scan(&resultId)
		result.Close()
		println("\033[0;96m", "[session] : there is allready a session with this username :")

		statementUpdate, err3 := db.Prepare("UPDATE session SET uuid = ? WHERE username = ?")
		if err3 != nil {
			fmt.Println("\033[1;31m", "[session] : error, impossible update session :", uuid)
			return
		}
		statementUpdate.Exec(uuid, user_name)

		return

	}

	statement, err := db.Prepare("INSERT INTO session (uuid, username) VALUES (?, ?)")

	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, can't insert into database")
		return
	}

	statement.Exec(uuid, user_name)

}

func DeleteSession(uuid string) {
	db := OpenDataBase()

	statement, err := db.Prepare("DELETE FROM session WHERE uuid = ?")

	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, deleting from database")
		return
	}
	statement.Exec(uuid)
}
