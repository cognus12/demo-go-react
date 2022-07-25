package vite

import (
	"errors"
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

type Vite struct {
	Main     *Chunck
	Env      string
	Platform string
}

var Chuncks ManifestMap = map[string]*Chunck{}

var v *Vite = &Vite{}

func NewVite(cfg *ViteConfig) (*Vite, error) {
	setConfigDefaults(cfg)

	chunks, err := parseManifest(cfg.ManifestPath)

	if err != nil {
		return nil, err
	}

	Chuncks = chunks

	mainChunck, ok := chunks[cfg.MainEntry]

	if !ok {
		return nil, errors.New("Wrong main chunk name")
	}

	v.Main = mainChunck
	v.Platform = cfg.Platform

	return v, nil
}
