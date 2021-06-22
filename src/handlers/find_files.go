package handlers

func findPathFiles(bodyPath string) []string {

	pathsOfFiles := []string{"./templates/layout.html", "./templates/includes/navbar.html", "./templates/includes/footer.html", "./templates/createpost.html", "./templates/createcomment.html"}
	pathsOfFiles = append(pathsOfFiles, bodyPath)

	return pathsOfFiles
}
