package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type ChunckMap = map[string]interface{}
type ManifestMap = map[string]ChunckMap

type Chunck struct {
	Key            string
	Src            string   `json:"src"`
	File           string   `json:"file"`
	Css            []string `json:"css"`
	Assets         []string `json:"assets"`
	IsEntry        bool     `json:"isEntry"`
	IsDynamicEntry bool     `json:"isDynamicEntry"`
	Imports        []string `json:"imports"`
	DynamicImports []string `json:"dynamicImports"`
}

func getManifest() {
	manifest, err := os.Open("static/manifest.json")

	if err != nil {
		log.Fatal("Failed to load manifest.json")
	}

	// var v interface{}

	var v ManifestMap

	byteValue, _ := ioutil.ReadAll(manifest)

	json.Unmarshal(byteValue, &v)

	processManifest(v)

	defer manifest.Close()
}

func logTarget(t []Chunck) {
	for _, v := range t {
		Printfln("---")
		Printfln("Key: %v", v.Key)
		Printfln("Src: %v", v.Src)
		Printfln("File: %v", v.File)
		Printfln("Css: %v", v.Css)
		Printfln("Assets: %v", v.Assets)
		Printfln("IsEntry: %v", v.IsEntry)
		Printfln("IsDynamicEntry: %v", v.IsDynamicEntry)
		Printfln("Imports: %v", v.Imports)
		Printfln("DynamicImports: %v", v.DynamicImports)
	}
}

func processManifest(m ManifestMap) []Chunck {
	target := []Chunck{}

	for k, v := range m {
		chunck := processChunck(v)

		chunck.Key = k

		target = append(target, chunck)
	}

	logTarget(target)

	return target
}

func processStringSlice(assets interface{}) []string {

	slice := []string{}

	v := reflect.ValueOf(assets)

	if v.Kind() == reflect.Slice {

		if (v.Len()) == 0 {
			return slice
		}

		for i := 0; i < v.Len(); i++ {
			slice = append(slice, fmt.Sprintf("%v", v.Index(i)))
		}
	}

	return slice
}

func processBool(v interface{}) bool {
	return reflect.ValueOf(v).Bool()
}

func processChunck(ch ChunckMap) Chunck {

	chunck := Chunck{}

	for k, v := range ch {

		if k == "src" {
			chunck.Src = fmt.Sprintf("%v", v)
		}

		if k == "file" {
			chunck.File = fmt.Sprintf("%v", v)
		}

		if k == "css" {
			chunck.Css = processStringSlice(v)
		}

		if k == "assets" {
			chunck.Assets = processStringSlice(v)
		}

		if k == "imports" {
			chunck.Imports = processStringSlice(v)
		}

		if k == "dynamicImports" {
			chunck.DynamicImports = processStringSlice(v)
		}

		if k == "isEntry" {
			chunck.IsEntry = processBool(v)
		}

		if k == "IsDynamicEntry" {
			chunck.IsDynamicEntry = processBool(v)
		}

	}

	return chunck
}
