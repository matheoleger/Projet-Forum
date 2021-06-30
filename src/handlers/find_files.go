package handlers

func findPathFiles(bodyPath string) []string {

	// On ajoute à un tableau, tout les fichiers html constituant la base d'une page
	pathsOfFiles := []string{"./templates/layout.html", "./templates/includes/navbar.html", "./templates/includes/footer.html", "./templates/createpost.html"}
	pathsOfFiles = append(pathsOfFiles, bodyPath)

	return pathsOfFiles
}
