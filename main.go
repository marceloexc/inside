package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type Page struct {
	Galleries []Gallery
}

type Gallery struct {
	Filepath string
	IndexSvg string
	Date     string
}

func compatibleGallery(entries []os.DirEntry) bool {
	hasHtml := false
	hasSvg := false
	hasDate := false

	//	check if the directory has both "index.html" and "index.svg"
	for _, entry := range entries {
		name := entry.Name()
		hasHtml = hasHtml || (name == "index.html")
		hasSvg = hasSvg || (name == "index.svg")
		hasDate = hasDate || (name == "date")
	}
	if hasHtml && hasSvg && hasDate {
		fmt.Println("Good!")
		return true
	}
	return false
}

func getGalleries(publicPath string) ([]Gallery, error) {

	var galleries []Gallery

	entries, err := os.ReadDir(publicPath)
	if err != nil {
		fmt.Println(err)
	}

	for _, e := range entries {
		if e.IsDir() {

			// from here i need to change cwd to public/...because then all of the links
			// to the pages go to ~/public/page instead of ~/page and it makes it all wrong.
			// or maybe not change cwd but just make this better, yanno?
			galleryPath := filepath.Join(publicPath, e.Name())

			galleryContents, _ := os.ReadDir(galleryPath)

			if compatibleGallery(galleryContents) {

				dateFilePath := filepath.Join(galleryPath, "date")
				dateContent, err := os.ReadFile(dateFilePath)
				if err != nil {
					fmt.Println("Couldn't read date")
					break
				}

				fileName := filepath.Base(galleryPath)

				galleries = append(galleries, Gallery{
					Filepath: fileName,
					IndexSvg: filepath.Join(fileName, "index.svg"),
					Date:     string(dateContent),
				})
			}
		}
	}
	return galleries, nil
}

func copyFiles(source string, destination string) {
	// can't figure out how to handle this error
	data, _ := os.ReadFile(source)
	_ = os.WriteFile(destination, data, 0644)
	fmt.Println("Copying file", source, " to ", destination)
}

func generatePage(templatePath string, publicPath string, page Page) {

	// copyFiles over assets to public folder
	templateFiles, err := os.ReadDir(templatePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range templateFiles {
		if !file.IsDir() {
			source := filepath.Join(templatePath, file.Name())
			destination := filepath.Join(publicPath, file.Name())
			copyFiles(source, destination)
		}
	}

	t, err := template.New("template.html").ParseFiles(filepath.Join(templatePath, "template.html"))

	f, err := os.Create(filepath.Join(publicPath, "index.html"))
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}

	if err != nil {
		fmt.Println("Couldn't make template")
	}
	err = t.Execute(f, page)

	//prints out template!
	err = f.Close()
	if err != nil {
		return
	}
}

func main() {

	absPath, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	fmt.Println(absPath)

	publicPath := filepath.Join("public")
	galleries, err := getGalleries(publicPath)

	if err != nil {
		fmt.Println("Could not find galleries")
		return
	}
	fmt.Println("Compatible galleries found", galleries)

	page := Page{Galleries: galleries}

	templatePath := filepath.Join("assets")

	generatePage(templatePath, publicPath, page)

}
