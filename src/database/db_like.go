package database

import (
	"fmt"
)

// func IsLiked(element string, username string, id int, isLiked bool) Like {
// 	db := OpenDataBase()
// 	defer db.Close()

// 	var like Like

// 	statement, err := db.Prepare("SELECT id_likes, isLiked FROM likes WHERE " + element + "= ? AND username = ?")

// 	if err != nil {
// 		fmt.Println("error prepare IsLiked : ", err)
// 		return like
// 	}

// 	result, errQuery := statement.Query(id, username)

// 	if errQuery != nil {
// 		fmt.Println("error query IsLiked : ", errQuery)
// 		return like
// 	}

// 	var idLikes int
// 	// var isLikedDataBase bool

// 	for result.Next() {
// 		// result.Scan(&idLikes, &isLikedDataBase)
// 		result.Scan(&idLikes, &like.isLiked)
// 	}

// 	if idLikes == 0 {
// 		like.LikeSate = false
// 	}

// 	// if idLikes == 0 {
// 	// 	AddLike(db, element, username, id, isLiked)
// 	// } else if isLikedDataBase == isLiked && idLikes != 0 {
// 	// 	DeleteLike(db, idLikes)
// 	// } else if isLikedDataBase != isLiked && idLikes != 0 {
// 	// 	ChangeValueLiked(db, idLikes, isLiked)
// 	// }

// 	fmt.Println(result)

// 	return like

// }

func IsLiked(element string, username string, id int) Like {
	db := OpenDataBase()
	defer db.Close()

	var like Like

	statement, err := db.Prepare("SELECT id_likes, isLiked FROM likes WHERE " + element + "= ? AND username = ?")

	if err != nil {
		fmt.Println("error prepare IsLiked : ", err)
		return like
	}

	result, errQuery := statement.Query(id, username)

	if errQuery != nil {
		fmt.Println("error query IsLiked : ", errQuery)
		return like
	}

	// var idLikes int
	// var isLikedDataBase bool

	for result.Next() {
		// result.Scan(&idLikes, &isLikedDataBase)
		result.Scan(&like.IdLike, &like.IsLiked)
	}

	if like.IdLike == 0 {
		like.LikeSate = false
	} else {
		like.LikeSate = true
	}

	// if idLikes == 0 {
	// 	AddLike(db, element, username, id, isLiked)
	// } else if isLikedDataBase == isLiked && idLikes != 0 {
	// 	DeleteLike(db, idLikes)
	// } else if isLikedDataBase != isLiked && idLikes != 0 {
	// 	ChangeValueLiked(db, idLikes, isLiked)
	// }

	fmt.Println(result)

	return like

}

func AddLike(element string, username string, id int, isLiked bool) {

	db := OpenDataBase()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO likes (username," + element + ",isLiked) VALUES ( ?, ?, ?)")
	if err != nil {
		fmt.Println("error prepare AddLike : ", err)
		return
	}

	statement.Exec(username, id, isLiked)

}

func DeleteLike(idLikes int) {

	db := OpenDataBase()
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM likes WHERE id_likes = ?")
	if err != nil {
		fmt.Println("error prepare DeleteLike : ", err)
		return
	}

	statement.Exec(idLikes)
}

func ChangeValueLiked(idLikes int, changeIsLiked bool) {

	db := OpenDataBase()
	defer db.Close()

	statement, err := db.Prepare("UPDATE likes SET isLiked = ? WHERE id_likes = ?")

	if err != nil {
		fmt.Println("error prepare ChangeValueLiked : ", err)
	}

	statement.Exec(changeIsLiked, idLikes)
}
