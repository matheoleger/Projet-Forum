package handlers

import (
	"database/sql"
	"fmt"
	"time"

	bdd "../database"

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
	db := OpenDataBase()

	statement, err := db.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		fmt.Println("error prepare")
		return
	}

	statement.Exec(user)

}

func GetElement(user, element string) string {
	db := OpenDataBase()

	statement, err := db.Prepare("SELECT " + element + " FROM user WHERE username = ?")
	result, err2 := statement.Query(user)

	if err != nil || err2 != nil {
		fmt.Println("Error query")
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

func GetPost() []bdd.Post {
	db := OpenDataBase()

	result, err := db.Query("SELECT * FROM post WHERE id_post NOT BETWEEN 1 AND 10")

	if err != nil {
		fmt.Println("error query")
	}

	defer result.Close()

	var post bdd.Post
	var Arraypost []bdd.Post

	for result.Next() {
		result.Scan(&post.Id_post, &post.Title, &post.Content, &post.Username, &post.Date, &post.Number_like)

		// fmt.Println(&post.Date)
		// fmt.Println(post.Id_post, post.Title, post.Username, post.Content, post.Date, post.Number_like, post.Liked)

		// On ajoute au tableau chaque post

		Arraypost = append(Arraypost, post)
	}
	err = result.Err()
	return Arraypost
}

func createCategory(name string) {
	db := OpenDataBase()
	statement, err := db.Prepare("INSERT INTO category (name) VALUES (?)")

	if err != nil {
		fmt.Println("error prepare createCategory")
		return
	}
	statement.Exec(name)

	defer db.Close()

}

func deleteCategory(name string) {

	db := OpenDataBase()
	statement, err := db.Prepare("DELETE FROM category WHERE name = ?")
	if err != nil {
		fmt.Println("error prepare ")
		return
	}
	statement.Exec(name)

	defer db.Close()

}

func InsertPost(title string, content string, username string, date_post time.Time, Number_like int) {

	db := OpenDataBase()

	statement, err := db.Prepare("INSERT INTO post (title, content, username, date_post, Number_like) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Println("error prepare InsertPost")
		return
	}

	statement.Exec(title, content, username, date_post, Number_like)

	defer db.Close()

}

func GetLastedID() int {
	db := OpenDataBase()

	// Selection du dernier id
	statement, err := db.Query("SELECT id_post FROM post ORDER BY id_post DESC LIMIT 1")

	if err != nil {
		fmt.Println("Error Query")
	}

	defer statement.Close()

	var id_post int

	for statement.Next() {
		statement.Scan(&id_post)
	}
	return id_post

}

func deletePost(id_post int) {
	db := OpenDataBase()

	statement, err := db.Prepare("DELETE FROM post WHERE id_post = ?")

	if err != nil {
		fmt.Println("error prepare deletePost")
		return
	}
	statement.Exec(id_post)

	defer db.Close()
}

func insertComment(content string, username string, post int) {
	db := OpenDataBase()

	statement, err := db.Prepare("INSERT INTO comment (content, username, post) VAlUES (?, ?, ?)")

	if err != nil {
		fmt.Println("error prepare ")
		return
	}
	statement.Exec(content, username, post)

	defer db.Close()
}

func deleteComment(id_comment int) {
	db := OpenDataBase()

	statement, err := db.Prepare("DELETE FROM comment WHERE id_comment = ?")

	if err != nil {
		fmt.Println("error prepare ")
		return
	}
	statement.Exec(id_comment)

	defer db.Close()
}

// func insertBridge(B_id_post int, B_id_category string) {
// 	db := OpenDataBase()

// 	statement, err := db.Prepare("INSERT INTO bridge (B_id_post, B_id_category) VAlUES (?, ?)")

// 	if err != nil {
// 		fmt.Println("error prepare ")
// 		return
// 	}
// 	statement.Exec(B_id_post, B_id_category)

// 	defer db.Close()
// }
