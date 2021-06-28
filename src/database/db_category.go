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

func GetCategory(per_page int, page int) []Category {
	db := OpenDataBase()
	prepare, err := db.Prepare("SELECT * from category LIMIT ? OFFSET ?")

	var CategoriesList []Category

	if err != nil {
		fmt.Println("error prepare GetCategory")
		return CategoriesList
	}

	// dapage := 0
	// if page != 0 {
	// 	dapage = page*per_page - 1
	// }

	result, err2 := prepare.Query(per_page, page)

	if err2 != nil {
		fmt.Println("error query GetComment in resultCat : ", err2)
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
