package handlers

import (
	"sort"
	"time"

	bdd "../database"
)

func FiltresLikeCroissant() []bdd.Post {

	// Récupération de tout les posts
	getPost := GetPost()

	// Trier de chaque élément du post en fonction de sa valeur
	sort.Slice(getPost, func(i, j int) bool {

		return getPost[i].Number_like > getPost[j].Number_like
	})

	// Renvoye les posts triés
	return getPost
}

func FiltresLikeDecroissant() []bdd.Post {
	// Même principe que pour la fonction ci-dessus
	getPost := GetPost()

	sort.Slice(getPost, func(i, j int) bool {
		return getPost[i].Number_like < getPost[j].Number_like
	})

	return getPost
}

func SortDate() []bdd.Post {

	// Récupération des posts
	getpost := GetPost()

	// Trie des posts en fonction de leur date

	sort.Slice(getpost, func(i, j int) bool {

		time1, _ := time.Parse("2014-11-12 11:45:26", getpost[i].Date)

		time2, _ := time.Parse("2014-11-12 11:45:26", getpost[j].Date)

		return time1.Before(time2)
	})

	// Renvoie des posts triés
	return getpost
}

func SortUserPost(username string) []bdd.Post {

	// Récupération de tout les posts
	getpost := GetPost()

	// On parcourt tout les posts afin de séléctionner uniquement ceux créé par l'utilisateur connecté
	for index := 0; index < len(getpost); index++ {
		if getpost[index].Username != username {
			getpost = append(getpost[:index], getpost[index+3:]...)
		}
	}

	// Renvoye les posts triés
	return getpost

}
