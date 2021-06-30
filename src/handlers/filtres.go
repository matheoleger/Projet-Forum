package handlers

import (
	"sort"
	"time"
	"unicode/utf8"

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

// Trier les catégorie
type CategorySort []bdd.Category

// Fonction qui renvoie la longueur de la catégories
func (categorie CategorySort) Len() int {
	return len(categorie)
}

// Fonction qui trie les catégories en fonction de leur lettres
func (categorie CategorySort) Less(i, j int) bool {
	iRune, _ := utf8.DecodeRuneInString(categorie[i].Name)
	jRune, _ := utf8.DecodeRuneInString(categorie[j].Name)
	return int32(iRune) < int32(jRune)
}

// Fonction qui change d'ordre les catégories dans le tableau
func (categorie CategorySort) Swap(i, j int) {
	categorie[i], categorie[j] = categorie[j], categorie[i]
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
