package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser(user string, pw string, mail string) {
	db, err := sql.Open("sqlite3", "../BDD/sqlForum")

	if err != nil {
		fmt.Println("error")
		return
	}

	statement, err := db.Prepare("INSERT INTO user (username, password, mail) VALUES (?, ?, ?)")
	statement.Exec(user, pw, mail)
}

func DeleteUser(user string) {
	db, err := sql.Open("sqlite3", "../../BDD/sqlForum")

	if err != nil {
		fmt.Println("error")
		return
	}

	statement, err := db.Prepare("DELETE FROM user WHERE username = ?")
	statement.Exec(user)
	// var password string
	// var email string
	// for result.Next() {
	// 	result.Scan(&password, &email)
	// 	/* Faire quelque chose avec cette ligne */
	// 	fmt.Println(password + " " + email)
	// }
}

func DataBase() {
	db, err := sql.Open("sqlite3", "./test-sqlite1")

	if err != nil {
		fmt.Println("error")
		return
	}

	result, err := db.Query("SELECT password, email FROM User WHERE username = \"JohnBibi\"")
	var password string
	var email string
	for result.Next() {
		result.Scan(&password, &email)
		/* Faire quelque chose avec cette ligne */
		fmt.Println(password + " " + email)
	}
}
