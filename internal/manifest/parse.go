package manifest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func Parse(path string) (*ManifestData, error) {
	manifest, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer manifest.Close()

	var v manifestMap

	byteValue, _ := ioutil.ReadAll(manifest)

	json.Unmarshal(byteValue, &v)

	data := processManifest(v)

	return &data, nil
}

func GetMainChunck(resources *ManifestData) *Chunck {
	for _, chunck := range *resources {
		// TODO simplify key
		if chunck.Key == "src/main.tsx" {
			return &chunck
		}
	}
	return nil
}

func processManifest(m manifestMap) ManifestData {
	target := []Chunck{}

	for k, v := range m {
		chunck := processChunck(v)

		chunck.Key = k

		target = append(target, chunck)
	}

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

func processChunck(ch chunckMap) Chunck {

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
