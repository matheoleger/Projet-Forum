package database

// func GetPagination(get_type string, per_page int, page int, from string) Page {

// 	var page_struct Page

// 	if get_type == "category" {
// 		var result_category []Category = GetCategory()
// 		page_struct.Categories = result_category

// 		println("\033[1;31m", "pls ignore", from)

// 		return page_struct

// 	} else if get_type == "post" {
// 		var result_post []Post = GetPostByCategory(from)
// 		page_struct.Posts = result_post

// 		return page_struct

// 	} else if get_type == "comment" {
// 		intFrom, err := strconv.Atoi(from)
// 		if err != nil {
// 			println("\033[1;31m", "from is not a integer")
// 			return page_struct
// 		}
// 		var result_comments []Comment = GetComments(intFrom, per_page, page)
// 		page_struct.Comments = result_comments

// 		return page_struct

// 	} else {
// 		println("\033[1;31m", "error")

// 		return page_struct

// 	}
// }

// func GetCategoryPagination(per_page int, page int) []Category {

// 	var result_category []Category = GetCategory()

// 	return result_category
// }

// func GetPostPagination(category string, per_page int, page int) []Post {

// 	var result_post []Post = GetPostByCategory(category)

// 	return result_post
// }

// func GetCommentPagination(post_id int, per_page int, page int) []Comment {

// 	var result_comments []Comment = GetComments(post_id, per_page, page)

// 	return result_comments
// }
