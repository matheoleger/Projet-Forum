package handlers

import (
	"database/sql"
	"fmt"
	"time"

	bdd "../database"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() *sql.DB {

	// Ouverture de la base
	db, err := sql.Open("sqlite3", "BDD/BDD_Finalv2.db")

	if err != nil {
		fmt.Println("\033[1;31m", "error open")
	}

	// Fonction indispensable pour la manipulation de la base de donnée
	return db
}

func AddUser(user string, pw string, mail string) {

	db := OpenDataBase()

	// Ajout d'un futur utilisateur
	statement, err := db.Prepare("INSERT INTO user (username, password, email) VALUES (?, ?, ?)")

	if err != nil {
		fmt.Println("error prepare")
		return
	}

	// Execution de l'ajout de l'utilisateur
	statement.Exec(user, pw, mail)

	defer db.Close()
}

func DeleteUser(user string) {
	db := OpenDataBase()

	// Suppression d'un utilisateur
	statement, err := db.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		fmt.Println("error prepare")
		return
	}

	// Exécution de la suppression d'un utilisateur
	statement.Exec(user)

}

func GetElement(user, element string) string {
	db := OpenDataBase()

	// Selection d'un élément de la table user
	statement, err := db.Prepare("SELECT " + element + " FROM user WHERE username = ?")
	result, err2 := statement.Query(user)

	if err != nil || err2 != nil {

		// Gestion d'erreur BDD
		fmt.Println("Error query")
		return "error query"
	}

	var password string

	// Parcourir toutes les colonnes de la table choisie
	for result.Next() {
		result.Scan(&password)
	}

	defer db.Close()

	return password
}

func GetPost() []bdd.Post {
	db := OpenDataBase()

	// Sélection de toutes les informations de tout les posts
	result, err := db.Query("SELECT * FROM post WHERE id_post NOT BETWEEN 1 AND 10")

	if err != nil {
		// Gestion d'erreur BDD
		fmt.Println("error query")
	}

	defer result.Close()

	var post bdd.Post
	var Arraypost []bdd.Post

	// Parcourir toutes les colonnes de chaque tables
	for result.Next() {
		result.Scan(&post.Id_post, &post.Title, &post.Content, &post.Username, &post.Date, &post.Number_like)

		// Ajout au tableau de chaque post

		Arraypost = append(Arraypost, post)
	}
	err = result.Err()
	return Arraypost
}

func createCategory(name string) {
	db := OpenDataBase()

	// Création d'une catégorie
	statement, err := db.Prepare("INSERT INTO category (name) VALUES (?)")

	if err != nil {
		// Gestion d'erreur BDD
		fmt.Println("error prepare createCategory")
		return
	}

	// Exécution de la création d'une catégorie
	statement.Exec(name)

	defer db.Close()

}

func deleteCategory(name string) {

	db := OpenDataBase()

	// Suppression d'une categorie
	statement, err := db.Prepare("DELETE FROM category WHERE name = ?")
	if err != nil {

		// Gestion d'erreur BDD
		fmt.Println("error prepare ")
		return
	}

	// Exécution de la suppresion d'une catégorie
	statement.Exec(name)

	defer db.Close()

}

func InsertPost(title string, content string, username string, date_post time.Time, Number_like int) {
	// Ouveture BDD
	db := OpenDataBase()

	// Insertion de tout les élément nécessaire pour la création d'un post
	statement, err := db.Prepare("INSERT INTO post (title, content, username, date_post, Number_like) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		// Gestion erreur BDD
		fmt.Println("error prepare InsertPost")
		return
	}

	// Exécution de l'insertion des éléments
	statement.Exec(title, content, username, date_post, Number_like)

	// Fermeture BDD
	defer db.Close()

}

func GetLastedID() int {
	db := OpenDataBase()

	// Selection du dernier id
	statement, err := db.Query("SELECT id_post FROM post ORDER BY id_post DESC LIMIT 1")

	if err != nil {
		// Gestion Erreur BDD
		fmt.Println("Error Query")
	}

	defer statement.Close()

	var id_post int

	// Parcourir tout les colonnes de la table
	for statement.Next() {
		statement.Scan(&id_post)
	}
	return id_post

}

func deletePost(id_post int) {
	db := OpenDataBase()

	// Suppression d'un post
	statement, err := db.Prepare("DELETE FROM post WHERE id_post = ?")

	if err != nil {
		// Gestion d'erreur BDD
		fmt.Println("error prepare deletePost")
		return
	}

	// Execution de la suppression d'un post
	statement.Exec(id_post)

	defer db.Close()
}

func insertComment(content string, username string, post int) {
	db := OpenDataBase()

	// Création de commentaire avec toutes les informations nécessaire
	statement, err := db.Prepare("INSERT INTO comment (content, username, post) VAlUES (?, ?, ?)")

	if err != nil {

		// Gestion erreur BDD
		fmt.Println("error prepare ")
		return
	}

	// Exécution de la création de commentaire
	statement.Exec(content, username, post)

	defer db.Close()
}

func deleteComment(id_comment int) {
	db := OpenDataBase()

	// Suppression d'un commentaire
	statement, err := db.Prepare("DELETE FROM comment WHERE id_comment = ?")

	if err != nil {
		// Gestion d'erreur BDD
		fmt.Println("error prepare ")
		return
	}

	// Exécution de la suppression d'un commentaire
	statement.Exec(id_comment)

	defer db.Close()
}
