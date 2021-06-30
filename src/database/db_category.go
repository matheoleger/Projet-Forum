package database

import (
	"database/sql"
	"fmt"
)

func OpenDataBase() *sql.DB {
	// Ouverture de la base de donnée
	db, err := sql.Open("sqlite3", "BDD/BDD_Finalv2.db")

	if err != nil {
		fmt.Println("error open")
	}
	return db
}

func GetCategory(per_page int, page int) []Category {
	db := OpenDataBase()

	// Selection de toutes les information de la table category
	prepare, err := db.Prepare("SELECT * from category LIMIT ? OFFSET ?")

	var CategoriesList []Category

	if err != nil {
		// Gestion d'erreur BDD
		fmt.Println("error prepare GetCategory")
		return CategoriesList
	}

	result, err2 := prepare.Query(per_page, page)

	if err2 != nil {
		// Gestion d'erreur BDD
		fmt.Println("error query GetComment in resultCat : ", err2)
		return CategoriesList
	}

	var nameCategory string

	// Parcourir toutes les colonnes de catégorie
	for result.Next() {
		result.Scan(&nameCategory)

		// Ajout de chaque catégorie dans un tableau
		CategoriesList = append(CategoriesList, Category{Name: nameCategory})
	}

	defer db.Close()

	// On renvoie la tableau contenant toutes les catégories
	return CategoriesList
}
