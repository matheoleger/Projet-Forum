package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser(user string, pw string, mail string) {
	db, err := sql.Open("sqlite3", "BDD/BDD_Forum")

	if err != nil {
		fmt.Println("error")
		return
	}

	statement, err := db.Prepare("INSERT INTO user (username, password, mail) VALUES (?, ?, ?)")

	if err != nil {
		fmt.Println("error")
		return
	}

	statement.Exec(user, pw, mail)
}

func DeleteUser(user string) {
	db, err := sql.Open("sqlite3", "BDD/BDD_Forum")

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
	db, err := sql.Open("sqlite3", "BDD/BDD_Forum")

	if err != nil {
		fmt.Println("error open")
		return
	}

	result, err := db.Query("SELECT password, mail FROM user WHERE username = \"JohnBibi\"")

	if err != nil {
		fmt.Println("error query")
		return
	}

	var password string
	var email string
	for result.Next() {
		result.Scan(&password, &email)
		/* Faire quelque chose avec cette ligne */
		fmt.Println(password + " " + email)
	}
}
