package database

import (
	"fmt"
	"strconv"
)

func GetComments(id_post int, per_page int, page int) []Comment {
	db := OpenDataBase()

	// var postStruct []Post
	var comments []Comment
	var comment Comment

	// statement, err := db.Prepare("SELECT id_comment, content, username, post FROM comment WHERE post = ?")
	statement, err := db.Prepare("SELECT id_comment, content, username, post FROM comment ORDER BY id_comment ASCENDING LIMIT " + strconv.Itoa(per_page) + " OFFSET " + strconv.Itoa(per_page*page))

	if err != nil {
		fmt.Println("error prepare GetComment in resultCat : ", err)
		return comments
	}

	result, err2 := statement.Query(id_post)

	if err2 != nil {
		fmt.Println("error query GetComment in resultCat : ", err2)
		return comments
	}

	// var title string
	// var content string
	// var username string

	for result.Next() {
		result.Scan(&comment.Id_comment, &comment.Content, &comment.Username, &comment.Post)

		comments = append(comments, comment)

		fmt.Println(comment)
	}

	return comments
}
