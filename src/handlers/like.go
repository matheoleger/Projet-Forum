package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	bdd "../database"
)

func Like(w http.ResponseWriter, r *http.Request) {
	postValue := r.URL.Query().Get("post")
	commentValue := r.URL.Query().Get("comment")
	isLikedValue := r.URL.Query().Get("isLiked")
	changeIsLiked, errConv := strconv.ParseBool(isLikedValue)

	if errConv != nil {
		return
	}

	if !VerifyCookie(w, r) {

		http.Redirect(w, r, "/login/", http.StatusSeeOther)

	} else {

		fmt.Println("post " + postValue + " comment " + commentValue + " isliked " + isLikedValue)

		username := bdd.GetProfil(w, r).Username

		var returningValue []byte

		if postValue != "" {

			returningValue = ChangeLike("post", postValue, changeIsLiked, username)

		} else if commentValue != "" {

			returningValue = ChangeLike("comment", commentValue, changeIsLiked, username)

		}

		w.Write(returningValue)
	}

}

func ChangeLike(elementType string, elementId string, changeIsLiked bool, username string) []byte {

	var like bdd.Like

	var finalResult []byte

	fmt.Println("there are postValue")
	elementIdInt, err := strconv.Atoi(elementId)

	if err != nil {
		return finalResult
	}

	like = bdd.IsLiked(elementType, username, elementIdInt)

	var nbrLike int

	if !like.LikeState {
		bdd.AddLike(elementType, username, elementIdInt, changeIsLiked)

		if changeIsLiked {
			nbrLike = bdd.ChangeNumberLike("+ 1", elementType, elementIdInt)
		} else if !changeIsLiked {
			nbrLike = bdd.ChangeNumberLike("- 1", elementType, elementIdInt)
		}

	} else if like.LikeState && like.IsLiked == changeIsLiked {
		bdd.DeleteLike(like.IdLike)

		if changeIsLiked {
			nbrLike = bdd.ChangeNumberLike("- 1", elementType, elementIdInt)
		} else if !changeIsLiked {
			nbrLike = bdd.ChangeNumberLike("+ 1", elementType, elementIdInt)
		}

	} else if like.LikeState && like.IsLiked != changeIsLiked {
		bdd.ChangeValueLiked(like.IdLike, changeIsLiked)

		if changeIsLiked {
			nbrLike = bdd.ChangeNumberLike("+ 2", elementType, elementIdInt)
		} else if !changeIsLiked {
			nbrLike = bdd.ChangeNumberLike("- 2", elementType, elementIdInt)
		}
	}

	var newLike = bdd.IsLiked(elementType, username, elementIdInt)

	var finalResultStruct bdd.ReturningLike

	finalResultStruct.LikeState = newLike.LikeState
	finalResultStruct.IsLiked = newLike.IsLiked
	finalResultStruct.Number_like = nbrLike

	finalResult, err = json.Marshal(finalResultStruct)

	if err != nil {
		fmt.Println("error json.Marshal returningLike : ", err)
	}

	return finalResult
}
