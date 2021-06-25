package database

import "strconv"

func GetPagination(get_type string, per_page int, page int, from string) Page {

	var page_struct Page

	if get_type == "category" {
		var result_category []Category = GetCategory()
		page_struct.Categories = result_category

		println("\033[1;31m", "pls ignore", from)

		return page_struct

	} else if get_type == "post" {
		var result_post []Post = GetPostByCategory(from)
		page_struct.Posts = result_post

		return page_struct

	} else if get_type == "comment" {
		var result_comments []Comment = GetComments(strconv.Atoi(from), per_page, page)
		page_struct.Comments = result_comments

		return page_struct
	}
}
