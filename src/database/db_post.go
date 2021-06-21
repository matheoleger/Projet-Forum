package database

import "fmt"

func GetPostByCategory(category string) []Post {
	var postStruct []Post

	db := OpenDataBase()

	statementCat, errCat := db.Prepare("SELECT B_id_post WHERE B_name_category = ?")

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

	for resultCat.Next() {

		statementCat.Close()

		resultCat.Scan(&postByCategory)

		statement, err := db.Prepare("SELECT id_post, title, content, username, date_post FROM post WHERE id_post = ?")

		if err != nil {
			fmt.Println("error prepare GetPostByCategory : ", err)
			return postStruct
		}

		result, err2 := statement.Query(postByCategory)

		if err2 != nil {
			fmt.Println("error query GetPostByCategory : ", err2)
			return postStruct
		}

		var id int
		var title string
		var content string
		var username string

		for result.Next() {
			result.Scan(&id, &title, &content, &username)

			postStruct = append(postStruct, Post{Id_post: id, Title: title, Content: content, Username: username})
		}
	}

	defer db.Close()

	return postStruct
}
