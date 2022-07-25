package vite

import (
	"errors"
	"html/template"
)

type Chunck struct {
	Src            string   `json:"src"`
	File           string   `json:"file"`
	Css            []string `json:"css"`
	Assets         []string `json:"assets"`
	IsEntry        bool     `json:"isEntry"`
	IsDynamicEntry bool     `json:"isDynamicEntry"`
	Imports        []string `json:"imports"`
	DynamicImports []string `json:"dynamicImports"`
}

type ManifestMap = map[string]*Chunck
type TemplateArgs = map[string]interface{}

type TemplateData struct {
	Args TemplateArgs
	*Chunck
}

type Vite struct {
	Manifest ManifestMap
	Main     string
}

type ViteConfig struct {
	// relative path, eg template/template.html
	TemplatePath string
	// relative path, eg static/assets
	AssetsPath string
	// relative path, eg static/manifest.json
	ManifestPath string
	// production or development, default - production
	Env string
	// react, vue, svelte, default - react
	Platform string
	// main chunk name, default - src/main.tsx
	MainChunk string
}

var v *Vite = &Vite{}
var t *template.Template

func NewVite(cfg *ViteConfig) (*Vite, error) {

	if cfg.Env == "" {
		cfg.Env = "production"
	}

	if cfg.Platform == "" {
		cfg.Platform = "react"
	}

	if cfg.MainChunk == "" {
		cfg.MainChunk = "src/main.tsx"
	}

	chunks, err := parseManifest(cfg.ManifestPath)

	if err != nil {
		return nil, err
	}

	v.Manifest = chunks

	_, ok := chunks[cfg.MainChunk]

	if !ok {
		return nil, errors.New("Wrong main chunk name")
	}

	v.Main = cfg.MainChunk

	return v, nil
}
