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

func processReflectedBool(v *reflect.Value) (target bool) {
	if (*v).Kind() == reflect.Bool {
		target = (*v).Bool()
	} else {
		target = false
	}

	return
}

func mapChunck(c reflect.Value, dist HTMLData) {
	for _, k := range c.MapKeys() {
		key := k.Convert(c.Type().Key())
		value := c.MapIndex(key).Elem()

		if !value.IsZero() {
			(dist)[key.String()] = c.MapIndex(key).Interface()
		}
	}
}

func detectEntry(c reflect.Value) bool {
	val := c.MapIndex(reflect.ValueOf("isEntry").Convert(c.Type().Key()))

	if !val.IsValid() {
		return false
	}

	if val.IsZero() {
		return false
	}

	val = val.Elem()

	if processReflectedBool(&val) {
		return true
	}

	return false
}

func mapManifest(m any) (HTMLData, error) {
	v := reflect.ValueOf(m)

	if !isMap(&v) {
		return nil, errors.New("Manifest should be a valid JSON, see https://vitejs.dev/guide/backend-integration.html")
	}

	target := HTMLData{}
	rest := []HTMLData{}
	keys := v.MapKeys()

	for _, k := range keys {
		key := k.Convert(v.Type().Key())
		chunck := v.MapIndex(key).Elem()
		var isEntry bool = detectEntry(chunck)

		if isEntry {
			mapChunck(chunck, target)
		} else {
			var node = make(HTMLData)
			mapChunck(chunck, node)
			rest = append(rest, node)
		}
	}

	target["nodes"] = rest

	return target, nil
}

func parseManifest(dist *fs.FS, path string) (HTMLData, error) {
	bytes, err := read(*dist, path)

	if err != nil {
		return nil, err
	}

	var jsonData any

	json.Unmarshal(bytes, &jsonData)

	t, err := mapManifest(jsonData)

	return t, err
}
