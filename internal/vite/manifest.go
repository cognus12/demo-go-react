package vite

import (
	"encoding/json"
	"errors"
	"io/fs"
	"reflect"
)

func read(fsys fs.FS, path string) ([]byte, error) {
	content, err := fs.ReadFile(fsys, path)

	return content, err
}

func isMap(v *reflect.Value) bool {
	if (*v).Kind() == reflect.Map {
		return true
	}

	return false
}

func mapReflectedStringSlice(v *reflect.Value) (target []string) {
	if (*v).Kind() != reflect.Slice {
		return target
	}

	if (*v).Len() != 0 {
		for i := 0; i < (*v).Len(); i++ {
			target = append(target, (*v).Index(i).Elem().String())
		}
	}

	return
}

func transformReflectedBool(v *reflect.Value) (target bool) {
	if (*v).Kind() == reflect.Bool {
		target = (*v).Bool()
	} else {
		target = false
	}

	return
}

func mapReflectedChunk(c reflect.Value) *Chunck {
	if !isMap(&c) {
		return nil
	}

	keys := c.MapKeys()
	target := Chunck{}

	for _, k := range keys {
		key := k.Convert(c.Type().Key())
		value := c.MapIndex(key).Elem()
		// fmt.Println("key :", key, " value:", value, "kind:", value.Kind())

		keyStr := key.String()

		if keyStr == "file" {
			target.File = value.String()
		}

		if keyStr == "src" {
			target.Src = value.String()
		}

		if keyStr == "isEntry" {
			target.IsEntry = transformReflectedBool(&value)
		}

		if keyStr == "isDynamicEntry" {
			target.IsDynamicEntry = transformReflectedBool(&value)
		}

		if keyStr == "css" {
			target.CSS = mapReflectedStringSlice(&value)
		}

		if keyStr == "assets" {
			target.Assets = mapReflectedStringSlice(&value)
		}

		if keyStr == "imports" {
			target.Imports = mapReflectedStringSlice(&value)
		}

		if keyStr == "dynamicImports" {
			target.DynamicImports = mapReflectedStringSlice(&value)
		}
	}

	return &target
}

func parseManifest(fs fs.FS, path string) (ManifestMap, error) {
	bytes, err := read(fs, path)

	if err != nil {
		return nil, err
	}

	var jsonData interface{}
	json.Unmarshal(bytes, &jsonData)
	reflectedManifest := reflect.ValueOf(jsonData)

	if !isMap(&reflectedManifest) {
		return nil, errors.New("Provided manifest is not valid json")
	}

	var target ManifestMap = map[string]*Chunck{}

	keys := reflectedManifest.MapKeys()

	for _, k := range keys {
		key := k.Convert(reflectedManifest.Type().Key())
		value := reflectedManifest.MapIndex(key).Elem()
		target[key.String()] = mapReflectedChunk(value)
	}

	return target, nil
}
