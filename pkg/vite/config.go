package vite

import (
	"html/template"
	"io/fs"
)

type ViteConfig struct {
	//
	FS fs.FS
	// production or development, default - production
	Env string
	// react, vue, svelte, default - react
	Platform string
	// path to frontend app folder relative to the root (default - frontend)
	ProjectDir string
	//
	SrcDir string
	// relative path to FrontendFolder, eg dist
	OutDir string
	//
	AssetsDir string
	// AssetsURLPrefix (/assets/ for prod, /src/ for dev)
	AssetsURLPrefix string
	//
	DevServerHost string
	//
	DevServerPort string

	//
	Template *template.Template
}

var defaults = map[string]string{
	"Env":      "production",
	"Platform": "react",
	"SrcDir":   "src",
	"OutDir":   "dist",

	"AssetsURLPrefixProd": "/assets/",
	"AssetsURLPrefixDev":  "/src/",

	"AssetsDir":     "assets",
	"DevServerHost": "localhost",
	"DevServerPort": "3000",
}

func (cfg *ViteConfig) setProdDefaults() {
	if cfg.AssetsURLPrefix == "" {
		cfg.AssetsURLPrefix = defaults["AssetsURLPrefixProd"]
	}
}

func (cfg *ViteConfig) setDevDefaults() {
	if cfg.DevServerHost == "" {
		cfg.DevServerHost = defaults["DevServerHost"]
	}

	if cfg.DevServerPort == "" {
		cfg.DevServerPort = defaults["DevServerPort"]
	}

	if cfg.AssetsURLPrefix == "" {
		cfg.AssetsURLPrefix = defaults["AssetsURLPrefixDev"]
	}
}

func (cfg *ViteConfig) setDefaults() {
	if cfg.Env == "" {
		cfg.Env = defaults["Env"]
	}

	if cfg.Platform == "" {
		cfg.Platform = defaults["Platform"]
	}

	if cfg.AssetsDir == "" {
		cfg.AssetsDir = defaults["AssetsDir"]
	}

	if cfg.Env == "production" {
		cfg.setProdDefaults()
	}

	if cfg.Env == "development" {
		cfg.setDevDefaults()
	}
}
