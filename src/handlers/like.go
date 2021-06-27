package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	bdd "../database"
)

func Like(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, r.Header.Get("Referer"), 302)

	fmt.Println("je test des choses et cest super cool")

	postValue := r.URL.Query().Get("post")
	commentValue := r.URL.Query().Get("comment")
	isLikedValue := r.URL.Query().Get("isLiked")
	changeIsLiked, errConv := strconv.ParseBool(isLikedValue)

	// var like bdd.Like

	if errConv != nil {
		return
	}

	fmt.Println("post" + postValue + "comment" + commentValue + "isliked" + isLikedValue)

	username := bdd.GetProfil(w, r).Username

	// var isLiked bool

	if postValue != "" {

		ChangeLike("post", postValue, changeIsLiked, username)

		// fmt.Println("there are postValue")
		// postValueInt, err := strconv.Atoi(postValue)

		// if err != nil {
		// 	return
		// }

		// like = bdd.IsLiked("post", username, postValueInt)

	} else if commentValue != "" {

		ChangeLike("comment", commentValue, changeIsLiked, username)

		// fmt.Println("there are commentValue")
		// commentValueInt, err2 := strconv.Atoi(commentValue)

		// if err2 != nil {
		// 	return
		// }

		// like = bdd.IsLiked("comment", username, commentValueInt)
	}

	// var t = template.Must(template.New("name").Parse("false"))

	// // page := bdd.Page{IsLiked: false}
	// t.Execute(w, nil)

}

func ChangeLike(elementType string, elementId string, changeIsLiked bool, username string) {

	var like bdd.Like

	fmt.Println("there are postValue")
	elementIdInt, err := strconv.Atoi(elementId)

	if err != nil {
		return
	}

	like = bdd.IsLiked(elementType, username, elementIdInt)

	if !like.LikeSate {
		bdd.AddLike(elementType, username, elementIdInt, changeIsLiked)
	} else if like.LikeSate && like.IsLiked == changeIsLiked {
		bdd.DeleteLike(like.IdLike)
	} else if like.LikeSate && like.IsLiked != changeIsLiked {
		bdd.ChangeValueLiked(like.IdLike, changeIsLiked)
	}
}
