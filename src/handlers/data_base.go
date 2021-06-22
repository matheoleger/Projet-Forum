package handlers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type PostStruct struct {
	Id_post     int
	Id_comment  int
	Title       string
	Content     string
	post        int
	Username    string
	Number_like int
	Liked       bool
	Date        time.Time
}

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

func GetPost() []PostStruct {
	db := OpenDataBase()

	result, err := db.Query("SELECT * FROM post WHERE id_post NOT BETWEEN 1 AND 9")

	if err != nil {
		fmt.Println("error query")
	}

	defer result.Close()

	var post PostStruct
	var Arraypost []PostStruct

	for result.Next() {
		result.Scan(&post.Id_post, &post.Title, &post.Content, &post.Username, &post.Number_like, &post.Liked, &post.Date)

		fmt.Println(&post.Date)
		// fmt.Println(post.Id_post, post.Title, post.Username, post.Content, post.Date, post.Number_like, post.Liked)

		// On ajoute au tableau chaque post
		Arraypost = append(Arraypost, post)
	}

	fmt.Println(Arraypost)
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

func InsertPost(title string, content string, username string, Number_like int, liked bool, date_post time.Time) {

	db := OpenDataBase()

	statement, err := db.Prepare("INSERT INTO post (title, content, username, Number_like, liked, date_post) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Println("error prepare InsertPost")
		return
	}

	statement.Exec(title, content, username, Number_like, liked, date_post)

	defer db.Close()
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
