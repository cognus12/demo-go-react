package vite

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
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
	// Chuncks ManifestMap

	MainEntry string
	CSS       []string
	Assets    []string

	Env             string
	Platform        string
	FS              fs.FS
	ProjectPath     string
	DistFolder      string
	AssetsPath      string
	AssetsDir       string
	AssetsURLPrefix string
}

var v *Vite

func NewVite(cfg *ViteConfig) (*Vite, error) {
	setConfigDefaults(cfg)

	v = &Vite{
		DistFolder: cfg.OutDir,
	}

	v.Platform = cfg.Platform
	v.ProjectPath = cfg.ProjectDir
	v.FS = os.DirFS(cfg.ProjectDir)
	v.AssetsURLPrefix = cfg.AssetsURLPrefix
	v.AssetsDir = cfg.AssetsDir

	chunks, err := parseManifest(v.FS, path.Join(cfg.OutDir, "manifest.json"))

	if err != nil {
		return nil, err
	}

	// v.Chuncks = chunks

	mainChunck, ok := chunks[cfg.MainEntry]

	if !ok {
		return nil, errors.New("Wrong main chunk name")
	}

	v.MainEntry = mainChunck.File
	v.CSS = mainChunck.Css

	v.AssetsPath = path.Join(v.ProjectPath, v.DistFolder, v.AssetsDir)

	fmt.Println(v.AssetsPath)

	return v, nil
}
