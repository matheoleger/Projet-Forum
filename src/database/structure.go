package database

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
	LikeInfo    Like
	Date        string
}

type Comment struct {
	Id_comment int
	Content    string
	Username   string
	Post       string
	LikeInfo   Like
}

type User struct {
	Username string
	Email    string
}

type Like struct {
	IdLike   int
	LikeSate bool
	IsLiked  bool
}
