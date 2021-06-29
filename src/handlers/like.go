package handlers

import (
	"encoding/json"
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

	if !VerifyCookie(w, r) {
		return
	}

	fmt.Println("post " + postValue + " comment " + commentValue + " isliked " + isLikedValue)

	username := bdd.GetProfil(w, r).Username

	// var isLiked bool

	var returningValue []byte

	if postValue != "" {

		returningValue = ChangeLike("post", postValue, changeIsLiked, username)

		// fmt.Println("there are postValue")
		// postValueInt, err := strconv.Atoi(postValue)

		// if err != nil {
		// 	return
		// }

		// like = bdd.IsLiked("post", username, postValueInt)

	} else if commentValue != "" {

		returningValue = ChangeLike("comment", commentValue, changeIsLiked, username)

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

	w.Write(returningValue)

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

	// if !like.LikeState && changeIsLiked {
	// 	bdd.ChangeNumberLike("+ 1", elementType, elementIdInt)
	// } else if !like.LikeState && !changeIsLiked {
	// 	bdd.ChangeNumberLike("- 1", elementType, elementIdInt)
	// } else if like.LikeState && changeIsLiked {
	// 	bdd.ChangeNumberLike("+ 2", elementType, elementIdInt)
	// } else if like.LikeState && !changeIsLiked {
	// 	bdd.ChangeNumberLike("- 2", elementType, elementIdInt)
	// }

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

	// finalResult := "{likeState: " + strconv.FormatBool(newLike.LikeState) + ", isLiked: " + strconv.FormatBool(newLike.IsLiked) + ", numberLike: " + strconv.Itoa(nbrLike) + "}"

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
