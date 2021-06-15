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

	statement, err := db.Prepare("INSERT INTO session (uuid, username) VALUES (?, ?)")

	//Error TO DO
	if err != nil {
		fmt.Println("\033[1;31m", "[session] : error, can't insert into database")
		return
	}
	println("\033[0;32m", "[session] : session sucessfully created : uuid = ", uuid, " username =", user_name)

	statement.Exec(uuid, user_name)
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
		fmt.Println("\033[1;31m", "[session] : error, cdelete insert into database")
		return
	}
	println("\033[0;32m", "[session] : session sucessfully deleted : uuid = ", uuid)

	statement.Exec(uuid)
}
