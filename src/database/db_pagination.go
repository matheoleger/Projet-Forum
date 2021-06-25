package database

func GetPagination(get_type string, per_page int, page int) {

	if get_type == "category" {

	} else if get_type == "post" {

	} else if get_type == "comment" {

	} else {
		println("\033[1;31m", "GetPagination error, type isn't category, post, or comment")
	}
}
