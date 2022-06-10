package web

import (
	"govite/internal/manifest"
	"html/template"
	"log"
	"os"
)

type ViewData struct {
	Title string
	*manifest.Chunck
}

var data *ViewData
var index *template.Template

func init() {
	mode := os.Getenv("MODE")

	log.Println("Run server in mode: ", mode)

	if mode == "api" {
		return
	}

	var resources *manifest.ManifestData
	var mainChunck *manifest.Chunck
	var manifestError error

	resources, manifestError = manifest.Parse("static/manifest.json")

	if manifestError != nil {
		log.Fatal("Failed to load manifest.json")
	}

	log.Println("manifest.json loaded")

	mainChunck = manifest.GetMain(resources)

	if mainChunck == nil {
		log.Fatal("Failed to load assets")
	}

	log.Println("Main chunck data parsed")

	data = &ViewData{Title: "Go+Vite App", Chunck: mainChunck}

	var templateErr error

	index, templateErr = template.ParseFiles("template.html")

	if templateErr != nil {
		log.Fatal("Template loading error: ", templateErr.Error())
	}

	log.Println("index.html template loaded")
}
