package database

import "time"

type Page struct {
	UserInfo   User
	Categories []Category
	Posts      []Post
	Comments   []Comment
}

type Category struct {
	Name string
}

type Post struct {
	Id_post     int
	Title       string
	Content     string
	Username    string
	Number_like int
	Liked       bool
	Date        time.Time
}

type Comment struct {
	Id_comment int
	Content    string
	Username   string
	post       string
	Liked      bool
}

type User struct {
	Username string
	Email    string
}
