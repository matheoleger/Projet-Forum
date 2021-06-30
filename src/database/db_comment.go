package database

import (
	"fmt"
)

func GetComments(id_post int, per_page int, page int) []Comment {
	// Ouverture base de donnée
	db := OpenDataBase()

	var comments []Comment
	var comment Comment

	// Selection de certains éléments de la table comment
	statement, err := db.Prepare("SELECT id_comment, content, username, post, Number_like FROM comment WHERE post = ? ORDER BY id_comment LIMIT ? OFFSET ?")

	if err != nil {
		fmt.Println("error prepare GetComment in resultCat : ", err)
		return comments
	}

	result, err2 := statement.Query(id_post, per_page, page)

	if err2 != nil {
		fmt.Println("error query GetComment in resultCat : ", err2)
		return comments
	}

	// Parcourir toute les colonnes selectionné de la table comment
	for result.Next() {
		result.Scan(&comment.Id_comment, &comment.Content, &comment.Username, &comment.Post, &comment.Number_like)

		// Ajouts des commentaires dans un tableau
		comments = append(comments, comment)

	}

	return comments
}
