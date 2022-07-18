package web

import (
	"demo-go-react/internal/manifest"
	"errors"
	"html/template"
	"log"
	"os"
	"sync"
)

type ViewParams = map[string]interface{}
type ViewData struct {
	Vars ViewParams
	*manifest.Chunck
}

var once sync.Once

var data *ViewData
var index *template.Template

const MODE_API = "api"
const MODE_FULL = "full"

func prepareResources() (*[]manifest.Chunck, error) {
	var resources *manifest.ManifestData
	var manifestError error
	resources, manifestError = manifest.Parse("static/manifest.json")
	return resources, manifestError
}

func prepareMainAssets(r *[]manifest.Chunck) (*manifest.Chunck, error) {
	var mainChunck *manifest.Chunck
	var err error
	mainChunck = manifest.GetMainChunck(r)
	if mainChunck == nil {
		err = errors.New("Failde to load main assets")
	}
	return mainChunck, err
}

func prepareData(v ViewParams, c *manifest.Chunck) {
	data = &ViewData{Vars: v, Chunck: c}
}

func prepareTemplate(v ViewParams) {
	resources, manifestError := prepareResources()

	if manifestError != nil {
		log.Fatal("Failed to load manifest.json")
	}

	log.Println("manifest.json loaded")

	mainChunck, mainChunckErr := prepareMainAssets(resources)

	if mainChunckErr != nil {
		log.Fatal(mainChunckErr)
	}

	log.Println("Main chunck data parsed")

	prepareData(v, mainChunck)

	var templateErr error
	index, templateErr = template.ParseFiles("template.html")

	if templateErr != nil {
		log.Fatal("Template loading error: ", templateErr.Error())
	}

	log.Println("index.html template loaded")
}

func Initialize(v ViewParams) {
	once.Do(func() {
		mode := os.Getenv("MODE")

		log.Println("Run server in mode: ", mode)

		if mode == MODE_API {
			return
		}

		prepareTemplate(v)
	})
}
