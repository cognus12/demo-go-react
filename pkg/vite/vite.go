package vite

import (
	"fmt"
	"html/template"
	"io/fs"
	"path"
)

type AssetsData = map[string]any
type Vite struct {
	FS          fs.FS
	DistFS      fs.FS
	Env         string
	Platform    string
	ProjectPath string

	SrcDir          string
	AssetsPath      string
	AssetsDir       string
	AssetsURLPrefix string

	MainEntryPath string
	DevServerURL  string

	data    AssetsData
	chuncks *[]AssetsData

	Template *template.Template
}

var v *Vite

func NewVite(cfg *ViteConfig) (*Vite, error) {
	cfg.setDefaults()

	v = &Vite{
		data: AssetsData{},
	}

	v.Env = cfg.Env
	v.Platform = cfg.Platform
	v.Template = cfg.Template

	distFs, err := fs.Sub(cfg.FS, "static")

	if err != nil {
		return nil, err
	}

	v.DistFS = distFs

	if v.Env == "production" {
		v.AssetsDir = cfg.AssetsDir

		err := v.parseManifest(&v.DistFS, "manifest.json")

		if err != nil {
			return nil, err
		}

		v.AssetsPath = path.Join("static", v.AssetsDir)
	}

	if v.Env == "development" {
		v.SrcDir = cfg.SrcDir
		v.DevServerURL = fmt.Sprintf("http://%v:%v", cfg.DevServerHost, cfg.DevServerPort)
	}

	v.AssetsURLPrefix = cfg.AssetsURLPrefix

	return v, nil
}
