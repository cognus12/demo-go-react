package vite

import (
	"fmt"
	"io/fs"
	"log"
	"path"
)

type HTMLData = map[string]any
type Vite struct {
	MainEntry string
	CSS       []string
	Assets    []string
	Imports   []string

	FS          fs.FS
	DistFS      fs.FS
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

	data HTMLData
}

var v *Vite

func NewVite(cfg *ViteConfig) (*Vite, error) {
	setConfigDefaults(cfg)

	v = &Vite{
		OutDir: cfg.OutDir,
	}

	v.Env = cfg.Env
	v.Platform = cfg.Platform
	v.ProjectPath = cfg.ProjectDir

	distFs, err := fs.Sub(cfg.FS, cfg.ProjectDir)

	if err != nil {
		log.Fatal(err)
	}

	v.DistFS = distFs

	v.MainEntryPath = fmt.Sprintf("%v/%v", cfg.SrcDir, cfg.MainEntry)

	if v.Env == "production" {

		v.AssetsDir = cfg.AssetsDir

		resources, err := parseManifest(&v.DistFS, path.Join(cfg.OutDir, "manifest.json"))

		if err != nil {
			return nil, err
		}

		v.data = resources

		v.AssetsPath = path.Join(v.ProjectPath, v.OutDir, v.AssetsDir)
	}

	if v.Env == "development" {
		v.SrcDir = cfg.SrcDir
		v.AssetsPath = path.Join(v.ProjectPath, v.SrcDir)
		v.DevServerURL = fmt.Sprintf("http://%v:%v", cfg.DevServerHost, cfg.DevServerPort)
	}

	v.AssetsURLPrefix = cfg.AssetsURLPrefix

	return v, nil
}
