package web

import (
	"govite/internal/manifest"
	"log"
)

type ViewData struct {
	Title string
	// Favicon string
	*manifest.Chunck
}

var resources *manifest.ManifestData
var mainChunck *manifest.Chunck

var data ViewData

func PrepareIndexPageData() *ViewData {
	var manifestError error

	resources, manifestError = manifest.Parse("static/manifest.json")

	if manifestError != nil {
		log.Fatal("Failed to load manifest.json")
	}

	mainChunck = manifest.GetMain(resources)

	if mainChunck == nil {
		log.Fatal("Failed to load assets")
	}

	data = ViewData{Title: "Go+Vite App", Chunck: *&mainChunck}

	return &data
}
