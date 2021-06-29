package handlers

import (
	"fmt"
	"sort"
	"unicode/utf8"

	bdd "../database"
)

func FiltresLike() []bdd.Post {
	getPost := GetPost()

	sort.Slice(getPost, func(i, j int) bool {
		return getPost[i].Number_like > getPost[j].Number_like
	})

	// fmt.Println(getPost)

	return getPost
}

// Trier les catégorie
type CategorySort []bdd.Category

func (categorie CategorySort) Len() int {
	return len(categorie)
}
func (categorie CategorySort) Less(i, j int) bool {
	iRune, _ := utf8.DecodeRuneInString(categorie[i].Name)
	jRune, _ := utf8.DecodeRuneInString(categorie[j].Name)
	return int32(iRune) < int32(jRune)
}
func (categorie CategorySort) Swap(i, j int) {
	categorie[i], categorie[j] = categorie[j], categorie[i]
}

func FiltresCategory() {
	test := bdd.GetCategory(20, 0)

	// fmt.Println(test)

	sort.Sort(CategorySort(test))

	for index := 0; index < len(test); index++ {
		bdd.GetPostByCategory(test[index].Name, 20, 0)

		fmt.Println("Category " + test[index].Name + " : ")
		fmt.Println(bdd.GetPostByCategory(test[index].Name, 20, 0))
	}

}
