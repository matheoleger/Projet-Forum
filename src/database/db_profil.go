package database

import (
	"fmt"
	"net/http"
)

func GetProfil(w http.ResponseWriter, r *http.Request) User {
	db := OpenDataBase()

	c, err := r.Cookie("session")

	if err != nil {
		fmt.Println("Cookie error")
	}
	test := c.Value

	statement, err := db.Prepare("SELECT username FROM session WHERE uuid = ?")

	if err != nil {
		fmt.Println("Error prepare")
	}

	result, err2 := statement.Query(test)

	if err2 != nil {
		fmt.Println("query error")
	}

	defer statement.Close()
	var username string
	for result.Next() {
		result.Scan(&username)
	}

	resultInformation := GetMoreInformation(w, r, username)
	return resultInformation
}

func GetMoreInformation(w http.ResponseWriter, r *http.Request, username string) User {
	db := OpenDataBase()

	var myUser User
	statement, err := db.Prepare("SELECT username, email FROM user WHERE username = ?")

	if err != nil {
		fmt.Println("Error prepare")
	}

	result, err2 := statement.Query(username)

	if err2 != nil {
		fmt.Println("query error")
	}

	defer statement.Close()

	for result.Next() {
		result.Scan(&myUser.Username, &myUser.Email)
		fmt.Println(myUser)
	}

	return myUser
}

func GetInformationAllUser(w http.ResponseWriter, r *http.Request, element string) []string {
	db := OpenDataBase()

	var users []string
	statement, err := db.Query("SELECT " + element + " FROM user")

	if err != nil {
		fmt.Println("Error prepares")
	}

	defer statement.Close()

	var test string
	for statement.Next() {
		statement.Scan(&test)
		users = append(users, test)
	}

	return users
}

func VerificationEmail(email string, emailbdd []string) bool {
	result := true
	for index := 0; index < len(emailbdd); index++ {
		if email == emailbdd[index] {
			result = false
			break
		}
	}
	return result
}
