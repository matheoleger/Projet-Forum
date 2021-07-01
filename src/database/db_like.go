package database

import (
	"fmt"
)

// Fonction qui récupère la colonne que l'on souhaite
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

	for result.Next() {
		result.Scan(&like.IdLike, &like.IsLiked)
	}

	if like.IdLike == 0 {
		like.LikeState = false
	} else {
		like.LikeState = true
	}

	fmt.Println(result)

	return like

}

// Fonction qui ajoute une ligne dans la table de like
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

// Fonction qui va retirer une ligne dans la table de like
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

// Fonction qui va changer la valeur isLiked pour le faire passer à sa valeur opposé
func ChangeValueLiked(idLikes int, changeIsLiked bool) {

	db := OpenDataBase()
	defer db.Close()

	statement, err := db.Prepare("UPDATE likes SET isLiked = ? WHERE id_likes = ?")

	if err != nil {
		fmt.Println("error prepare ChangeValueLiked : ", err)
	}

	statement.Exec(changeIsLiked, idLikes)
}

// Fonction qui va changer le nombre de like dans la BDD
func ChangeNumberLike(incSign string, elementType string, id int) int {
	db := OpenDataBase()
	defer db.Close()

	statement, err := db.Prepare("UPDATE " + elementType + " SET Number_like = Number_like " + incSign + " WHERE id_" + elementType + "= ? RETURNING Number_like ")

	if err != nil {
		fmt.Println("error prepare changeNumberLike : ", err)
	}

	var nbrLike int

	statement.QueryRow(id).Scan(&nbrLike)

	return nbrLike
}
