package vite

import (
	"errors"
	"io/fs"
	"os"
	"path"
)

type Chunck struct {
	Src            string   `json:"src"`
	File           string   `json:"file"`
	CSS            []string `json:"css"`
	Assets         []string `json:"assets"`
	IsEntry        bool     `json:"isEntry"`
	IsDynamicEntry bool     `json:"isDynamicEntry"`
	Imports        []string `json:"imports"`
	DynamicImports []string `json:"dynamicImports"`
}

type ManifestMap = map[string]*Chunck
type TemplateArgs = map[string]interface{}

type Vite struct {
	MainEntry string
	CSS       []string
	Assets    []string
	Imports   []string

	FS              fs.FS
	Env             string
	Platform        string
	ProjectPath     string
	OutDir          string
	AssetsPath      string
	AssetsDir       string
	AssetsURLPrefix string
	Data            map[string]any
}

var v *Vite

func NewVite(cfg *ViteConfig, data map[string]any) (*Vite, error) {
	setConfigDefaults(cfg)

	v = &Vite{
		OutDir: cfg.OutDir,
		Data:   data,
	}

	v.Env = cfg.Env
	v.Platform = cfg.Platform
	v.ProjectPath = cfg.ProjectDir
	v.FS = os.DirFS(cfg.ProjectDir)
	v.AssetsURLPrefix = cfg.AssetsURLPrefix
	v.AssetsDir = cfg.AssetsDir

	chunks, err := parseManifest(v.FS, path.Join(cfg.OutDir, "manifest.json"))

	if err != nil {
		return nil, err
	}

	mainChunck, ok := chunks[cfg.MainEntry]

	if !ok {
		return nil, errors.New("Wrong main chunk name")
	}

	v.MainEntry = mainChunck.File
	v.CSS = mainChunck.CSS
	v.Imports = mainChunck.Imports

	v.AssetsPath = path.Join(v.ProjectPath, v.OutDir, v.AssetsDir)

	return v, nil
}
