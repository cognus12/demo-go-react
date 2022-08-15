package vite

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
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

func processReflectedBool(v *reflect.Value) (target bool) {
	if (*v).Kind() == reflect.Bool {
		target = (*v).Bool()
	} else {
		target = false
	}

	return
}

func processReflectedString(v *reflect.Value) (target string) {
	target = (*v).String()

	return
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

func mapReflectedChunk(c reflect.Value) *Chunck {
	if !isMap(&c) {
		return nil
	}

	keys := c.MapKeys()
	target := Chunck{}

	for _, k := range keys {
		kk := k.Convert(c.Type().Key())
		value := c.MapIndex(kk).Elem()

		key := kk.String()

		if key == "file" {
			target.File = processReflectedString(&value)
		}

		if key == "src" {
			target.Src = processReflectedString(&value)
		}

		if key == "isEntry" {
			target.IsEntry = processReflectedBool(&value)
		}

		if key == "isDynamicEntry" {
			target.IsDynamicEntry = processReflectedBool(&value)
		}

		if key == "css" {
			target.CSS = mapReflectedStringSlice(&value)
		}

		if key == "assets" {
			target.Assets = mapReflectedStringSlice(&value)
		}

		if key == "imports" {
			target.Imports = mapReflectedStringSlice(&value)
		}

		if key == "dynamicImports" {
			target.DynamicImports = mapReflectedStringSlice(&value)
		}
	}

	return &target
}

func parseManifest(dist *fs.FS, path string) (ManifestMap, error) {
	bytes, err := read(*dist, path)

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

func printMap(m map[string]interface{}) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

type mtarget = map[string]interface{}

func mapChunck(c reflect.Value, dist mtarget) {
	for _, k := range c.MapKeys() {
		key := k.Convert(c.Type().Key())
		value := c.MapIndex(key).Elem()

		if !value.IsZero() {
			(dist)[key.String()] = c.MapIndex(key).Interface()
		}
	}
}

func mapManifest(m reflect.Value) mtarget {
	target := mtarget{}
	rest := []mtarget{}
	keys := m.MapKeys()

	for _, k := range keys {
		key := k.Convert(m.Type().Key())
		chunck := m.MapIndex(key).Elem()
		var isEntry bool = false

		for _, kk := range chunck.MapKeys() {
			key := kk.Convert(chunck.Type().Key())
			value := chunck.MapIndex(key).Elem()
			if key.String() == "isEntry" {
				val := processReflectedBool(&value)
				if val {
					isEntry = true
				}
			}
		}

		if isEntry {
			mapChunck(chunck, target)
		} else {
			var node = make(mtarget)
			mapChunck(chunck, node)
			rest = append(rest, node)
		}
	}

	target["nodes"] = rest

	return target
}

func parseManifestNew(dist *fs.FS, path string) (mtarget, error) {
	bytes, err := read(*dist, path)

	if err != nil {
		return nil, err
	}

	var jsonData interface{}
	json.Unmarshal(bytes, &jsonData)

	v := reflect.ValueOf(jsonData)

	if !isMap(&v) {
		log.Fatal("Manifest should be a valid JSON, see https://vitejs.dev/guide/backend-integration.html")
	}

	t := mapManifest(v)

	fmt.Println(t)

	return t, nil
}
