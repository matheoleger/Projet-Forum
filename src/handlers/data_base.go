package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser(user string, pw string, mail string) {
	db, err := sql.Open("sqlite3", "BDD/BBD_v5")

	if err != nil {
		fmt.Println("error open")
		return
	}

	statement, err := db.Prepare("INSERT INTO user (id_username, password, email) VALUES (?, ?, ?)")

	//Error TO DO
	if err != nil {
		fmt.Println("error prepare")
		return
	}
	statement.Exec(user, pw, mail)
}

func DeleteUser(user string) {
	db, err := sql.Open("sqlite3", "BDD/BBD_v5")

	if err != nil {
		fmt.Println("error open 1")
		return
	}

	statement, err := db.Prepare("DELETE FROM user WHERE id_username = ?")
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

func GetPassWord(user string) string {
	db, err := sql.Open("sqlite3", "BDD/BBD_v5")

	if err != nil {
		fmt.Println("error open")
		return "error open3"
	}

	statement, err := db.Prepare("SELECT password FROM user WHERE id_username = ?")
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

	return password
}
