package database

import (
	"database/sql"
	"fmt"
	// handlers "../handlers"
)

func OpenDataBase() *sql.DB {
	db, err := sql.Open("sqlite3", "BDD/BDD_Finalv2.db")

	if err != nil {
		fmt.Println("error open")
	}
	return db
}

func GetCategory() []Category {
	db := OpenDataBase()
	result, err := db.Query("SELECT * from category")

	var CategoriesList []Category

	if err != nil {
		fmt.Println("error query GetCategory")
		return CategoriesList
	}

	var nameCategory string
	for result.Next() {
		result.Scan(&nameCategory)
		CategoriesList = append(CategoriesList, Category{Name: nameCategory})
	}

	fmt.Println(CategoriesList)

	defer db.Close()

	return CategoriesList
}
