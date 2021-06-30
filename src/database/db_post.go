package database

import (
	"database/sql"
	"fmt"
	"time"
)

func GetPostByCategory(category string, per_page int, page int) []Post {
	var postStruct []Post

	// Ouverture BDD
	db := OpenDataBase()

	// Sélection de l'Id_post dans la table bridge afin de récupérer tout les posts de chaque catégorie
	statementCat, errCat := db.Prepare("SELECT B_id_post FROM bridge WHERE B_name_category = ? ORDER BY id LIMIT ? OFFSET ?")

	if errCat != nil {
		// Gestion erreur BDD
		fmt.Println("error prepare GetPostByCategory : ", errCat)
		return postStruct
	}

	resultCat, errQueryCat := statementCat.Query(category, per_page, page*per_page)

	if errQueryCat != nil {
		// Gestion erreur BDD
		fmt.Println("error prepare GetPostByCategory : ", errQueryCat)
		return postStruct
	}

	var postByCategory int

	statementCat.Close()

	// Parcourir les colonnes sélectionné de la table choisie
	for resultCat.Next() {

		resultCat.Scan(&postByCategory)

		// Ajout de chaque post en fonction de sa catégorie dans un tableau
		postStruct = append(postStruct, GetPost(db, postByCategory))

	}

	defer db.Close()

	return postStruct
}

func GetPost(db *sql.DB, id_post int) Post {

	// var postStruct []Post
	var post Post

	// Sélection de certains élément de la table post
	statement, err := db.Prepare("SELECT title, content, username, date_post, Number_like FROM post WHERE id_post = ?")

	if err != nil {
		// Gestion d'erreur BDD
		fmt.Println("error prepare GetPostByCategory in resultCat : ", err)
		return post
	}

	result, err2 := statement.Query(id_post)

	if err2 != nil {
		// Gestion d'erreur BDD
		fmt.Println("error query GetPostByCategory in resultCat : ", err2)
		return post
	}

	var date time.Time

	for result.Next() {
		result.Scan(&post.Title, &post.Content, &post.Username, &date, &post.Number_like)

		post.Id_post = id_post
	}

	// Change le format de la date, pour une meilleur lisibilité
	dateFormated := date.Format("2006-01-02 15:04:05")

	post.Date = dateFormated

	return post
}
