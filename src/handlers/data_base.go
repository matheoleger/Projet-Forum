package handlers

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() *sql.DB {
	db, err := sql.Open("sqlite3", "BDD/BDD_Finalv2.db")

	if err != nil {
		fmt.Println("error open")
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

func GetPost() interface{} {
	db := OpenDataBase()

	type PostStruct struct {
		Id_post     int
		post		string
		Title       string
		Content     string
		Username    string
		Number_like int
		Liked       bool
		Date        time.Time
	}

	//statement, err := db.Prepare("SELECT title, content, username FROM post WHERE id_post=$1")

	statement := db.QueryRow("SELECT id_post, title, content, username, Number_like, liked ,date_post FROM post WHERE id_post=8;", 4)

	var post PostStruct
	switch err := statement.Scan(&post.Id_post, &post.Title, &post.Content, &post.Username, &post.Number_like, &post.Liked, &post.Date); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("C'est le post numéro " + strconv.Itoa(post.Id_post) + " le titre est : " + post.Title + " le contenu du post est : " + post.Content + " c'est " + post.Username + " qui a écrit le post, et il y a  " + strconv.Itoa(post.Number_like) + " likes, et la date du post est " + post.Date.String())
	default:
		panic(err)
	}
	// statement.Scan(&title, &content, &username)
	return post
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

func InsertPost(title string, content string, username string) {
	db := OpenDataBase()

	statement, err := db.Prepare("INSERT INTO post (title, content, username) VAlUES (?, ?, ?)")

	if err != nil {
		fmt.Println("error prepare InsertPost")
		return
	}

	statement.Exec(title, content, username)

	defer db.Close()
}

func deletePost(id_post int) {
	db := OpenDataBase()

	statement, err := db.Prepare("DELETE FROM post WHERE id_post = ?")

	if err != nil {
		fmt.Println("error prepare deletePost")
		return
	}

	defer db.Close()
}

func insertComment(content string, username string,post string) {
	db := OpenDataBase()

	statement, err := db.Prepare("INSERT INTO comment (content, username, post) VAlUES (?, ?, ?)")

	if err != nil {
		fmt.Println("error prepare ")
		return
	}
	statement.Exec(content, username, post)

	defer db.Close()
}
