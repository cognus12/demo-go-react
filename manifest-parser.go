package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type Chunck struct {
	Src            string
	File           string
	Css            []string
	Assets         []string
	IsEntry        bool
	IsDynamicEntry bool
	Imports        []string
	DynamicImports []string
}

type manifestNode = map[string]Chunck

// type ManifestData struct {
// 	File string
// 	CSS  []string
// }

func getManifest() {
	manifest, err := os.Open("static/manifest.json")

	if err != nil {
		log.Fatal("Failed to load manifest.json")
	}

	var v interface{}

	byteValue, _ := ioutil.ReadAll(manifest)

	json.Unmarshal(byteValue, &v)

	processManifest(v)

	defer manifest.Close()
}

func processManifest(m interface{}) {
	v := reflect.ValueOf(m)

	switch v.Kind() {
	case reflect.Map:

		log.Printf("Loaded manifest raw data: %v", v.Interface())

		// for k, v := range v.Interface() {

		// }

		break
	default:
		log.Fatal("Incorrect manifest.json")
	}

}

/*

"file": "assets/main.4889e940.js",
    "src": "main.js",
    "isEntry": true,
    "dynamicImports": ["views/foo.js"],
    "css": ["assets/main.b82dbe22.css"],
    "assets": ["assets/asset.0ab0f9cd.png"]



	export interface ManifestChunk {
  src?: string
  file: string
  css?: string[]
  assets?: string[]
  isEntry?: boolean
  isDynamicEntry?: boolean
  imports?: string[]
  dynamicImports?: string[]
}


*/
