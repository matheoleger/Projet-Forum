package database

import (
	"database/sql"
	"fmt"
	"time"
)

func GetPostByCategory(category string, username string) []Post {
	var postStruct []Post

	db := OpenDataBase()

	statementCat, errCat := db.Prepare("SELECT B_id_post FROM bridge WHERE B_name_category = ?")

	if errCat != nil {
		fmt.Println("error prepare GetPostByCategory : ", errCat)
		return postStruct
	}

	resultCat, errQueryCat := statementCat.Query(category)

	if errQueryCat != nil {
		fmt.Println("error prepare GetPostByCategory : ", errQueryCat)
		return postStruct
	}

	var postByCategory int

	statementCat.Close()

	for resultCat.Next() {

		resultCat.Scan(&postByCategory)

		postStruct = append(postStruct, GetPost(db, postByCategory, username))

	}

	defer db.Close()

	return postStruct
}

func GetPost(db *sql.DB, id_post int, username string) Post {

	// var postStruct []Post
	var post Post

	statement, err := db.Prepare("SELECT title, content, username, date_post FROM post WHERE id_post = ?")

	if err != nil {
		fmt.Println("error prepare GetPostByCategory in resultCat : ", err)
		return post
	}

	result, err2 := statement.Query(id_post)

	if err2 != nil {
		fmt.Println("error query GetPostByCategory in resultCat : ", err2)
		return post
	}

	// var title string
	// var content string
	// var username string

	var date time.Time

	for result.Next() {
		result.Scan(&post.Title, &post.Content, &post.Username, &date)

		post.Id_post = id_post
		// postStruct = append(postStruct, Post{Id_post: id_post, Title: title, Content: content, Username: username})

		fmt.Println(post)
	}

	dateFormated := date.Format("2006-01-02 15:04:05")

	post.Date = dateFormated

	post.LikeInfo = IsLiked("post", username, id_post)

	fmt.Println(post.Date)

	return post
}
