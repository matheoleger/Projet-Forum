package database

type Page struct {
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
	Date        string
}

type Comment struct {
	Id_comment int
	Content    string
	Username   string
	post       string
	Liked      bool
}
