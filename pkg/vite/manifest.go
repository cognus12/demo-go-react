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

func mapChunck(c map[string]any, dist AssetsData) {

	v := reflect.ValueOf(c)

	for _, k := range v.MapKeys() {
		key := k.Convert(v.Type().Key())
		value := v.MapIndex(key).Elem()

		if !value.IsZero() {
			(dist)[key.String()] = v.MapIndex(key).Interface()
		}
	}
}

func mapManifest(m any) (AssetsData, []AssetsData, error) {
	manifest, ok := m.(map[string]any)

	if !ok {
		return nil, nil, errors.New("Manifest should be a valid JSON, see https://vitejs.dev/guide/backend-integration.html")
	}

	raw := AssetsData{}
	chuncks := []AssetsData{}

	for _, chunck := range manifest {
		m, ok := chunck.(map[string]any)

		if ok {
			isEntry, ok := m["isEntry"].(bool)

			if ok && isEntry {
				mapChunck(m, raw)
			} else {
				var node = make(AssetsData)
				mapChunck(m, node)
				chuncks = append(chuncks, node)
			}
		}

	}

	target := map[string]any{}

	target["file"] = raw["file"]
	target["css"] = raw["css"]
	target["assets"] = raw["assets"]
	target["imports"] = raw["imports"]

	return target, chuncks, nil
}

func parseManifest(dist *fs.FS, path string) (AssetsData, []AssetsData, error) {
	bytes, err := read(*dist, path)

	if err != nil {
		return nil, nil, err
	}

	var jsonData any

	json.Unmarshal(bytes, &jsonData)

	t, chuncks, err := mapManifest(jsonData)

	return t, chuncks, err
}
