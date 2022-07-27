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

	FS          fs.FS
	Env         string
	Platform    string
	ProjectPath string

	SrcDir          string
	OutDir          string
	AssetsPath      string
	AssetsDir       string
	AssetsURLPrefix string

	MainEntryPath string
	DevServerURL  string

	Data map[string]any
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

	v.MainEntryPath = fmt.Sprintf("%v/%v", cfg.SrcDir, cfg.MainEntry)

	if v.Env == "production" {
		v.AssetsURLPrefix = cfg.AssetsURLPrefix
		v.AssetsDir = cfg.AssetsDir

		chunks, err := parseManifest(v.FS, path.Join(cfg.OutDir, "manifest.json"))

		if err != nil {
			return nil, err
		}

		mainChunck, ok := chunks[v.MainEntryPath]

		if !ok {
			return nil, errors.New("Wrong main chunk name")
		}

		v.MainEntry = mainChunck.File
		v.CSS = mainChunck.CSS
		v.Imports = mainChunck.Imports

		v.AssetsPath = path.Join(v.ProjectPath, v.OutDir, v.AssetsDir)
	}

	if v.Env == "development" {
		v.SrcDir = cfg.SrcDir
		v.DevServerURL = fmt.Sprintf("http://%v:%v", cfg.DevServerHost, cfg.DevServerPort)
	}

	return v, nil
}
