package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() *sql.DB {
	db, err := sql.Open("sqlite3", "BDD/BDD_Finalv2.db")

	if err != nil {
		fmt.Println("\033[1;31m", "error open")
	}
	return db
}

func AddUser(user string, pw string, mail string) {

	db := OpenDataBase()
	statement, err := db.Prepare("INSERT INTO user (username, password, email) VALUES (?, ?, ?)")

	//Error TO DO
	if err != nil {
		fmt.Println("error prepare")
		return
	}
	statement.Exec(user, pw, mail)

	defer db.Close()
}

func DeleteUser(user string) {
	// db, err := sql.Open("sqlite3", "BDD/BBD_Final")

	// if err != nil {
	// 	fmt.Println("error open 1")
	// 	return
	// }
	//

	db := OpenDataBase()

	statement, err := db.Prepare("DELETE FROM user WHERE id_username = ?")
	if err != nil {
		fmt.Println("error prepare")
		return
	}

	statement.Exec(user)
	// var password string
	// var email string
	// for result.Next() {
	// 	result.Scan(&password, &email)
	// 	/* Faire quelque chose avec cette ligne */
	// 	fmt.Println(password + " " + email)
	// }
}

// func DataBase() {
// 	db, err := sql.Open("sqlite3", "BDD/BBD_v5")

// 	if err != nil {
// 		fmt.Println("error open2")
// 		return
// 	}

// 	result, err := db.Query("SELECT password, email FROM user WHERE id_username = \"JohnBibi\"")

// 	// Error TO DO
// 	if err != nil {
// 		fmt.Println("error query data base")
// 		return
// 	}

// 	var password string
// 	var email string
// 	for result.Next() {
// 		result.Scan(&password, &email)
// 		/* Faire quelque chose avec cette ligne */
// 		fmt.Println(password + " " + email)
// 	}
// }

func GetElement(user, element string) string {
	db := OpenDataBase()

	statement, err := db.Prepare("SELECT " + element + " FROM user WHERE username = ?")
	result, err2 := statement.Query(user)

	if err != nil || err2 != nil {
		fmt.Println("error query")
		return "error query"
	}

	var password string

	for result.Next() {
		result.Scan(&password)
		/* Faire quelque chose avec cette ligne */
		fmt.Println(password)
	}

	defer db.Close()

	return password
}

//	---------- uuid ----------

func AddSession(uuid string, user_name string) {
	db := OpenDataBase()

	// if err != nil {
	// 	fmt.Println("\033[1;31m", "[session] : error open")
	// 	return
	// }

	statementTest, err := db.Prepare("SELECT uuid FROM session WHERE username = ?")
	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, impossible to verify if session allready exist", uuid)
		return
	}
	// statementTest.Exec(user_name)
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

		println("\033[0;32m", "[session] : session sucessfully updated with uuid = ", uuid)

		return

	}

	statementuuid, err4 := db.Prepare("SELECT username FROM session WHERE uuid = ?")
	if err4 != nil {
		fmt.Println("\033[1;31m", "[session] : error, impossible to do statement")
		return
	}
	resultuuid, err5 := statementuuid.Query(uuid)
	if err5 != nil {
		fmt.Println("\033[1;31m", "[session] : error, impossible to querry")
		return
	}

	if resultId == uuid {
		fmt.Println("\033[1;31m", "[session] : error, a session allready exist with uuid =", uuid)

		return

		// } else if  == user_name {

		// 	println("\033[0;96m", "[session] : there is allready a session with this username :", err)

		// 	statementMega, err3 := db.Prepare("UPDATE session SET uuid = (uuid) WHERE username = (username) (?, ?)")
		// 	if err3 != nil {
		// 		fmt.Println("\033[1;31m", "[session] : error, impossible update session :", uuid)
		// 		return
		// 	}
		// 	statementMega.Exec(uuid, user_name)

		// 	println("\033[0;32m", "[session] : session sucessfully updated with uuid = ", uuid)

		// 	return

	} else {

		statement, err := db.Prepare("INSERT INTO session (uuid, username) VALUES (?, ?)")

		//Error TO DO
		if err != nil {
			fmt.Println("\033[1;31m", "[session] : error, can't insert into database")
			return
		}
		println("\033[0;32m", "[session] : session sucessfully created : uuid = ", uuid, " username =", user_name)

		statement.Exec(uuid, user_name)
	}
}

func DeleteSession(uuid string) {
	db := OpenDataBase()

	// if err != nil {
	// 	fmt.Println("\033[1;31m", "[session] : error open")
	// 	return
	// }

	statement, err := db.Prepare("DELETE FROM session WHERE uuid = ?")

	//Error TO DO
	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, deleting from database")
		return
	}
	println("\033[0;32m", "[session] : session sucessfully deleted : uuid = ", uuid)

	statement.Exec(uuid)
}
