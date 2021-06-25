package handlers

import (
	"fmt"
	"sort"

	bdd "../database"
)

func FiltresLike() []bdd.Post {
	getPost := GetPost()

	sort.Slice(getPost, func(i, j int) bool {
		return getPost[i].Number_like > getPost[j].Number_like
	})

	fmt.Println(getPost)

	return getPost
}
