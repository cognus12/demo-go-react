package vite

type ViteConfig struct {
	// production or development, default - production
	Env string
	// react, vue, svelte, default - react
	Platform string
	// main chunk name, default - src/main.tsx
	MainEntry string
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
}

var defaults = map[string]string{
	"Env":             "production",
	"Platform":        "react",
	"MainEntry":       "main.tsx",
	"AssetsURLPrefix": "/assets/",
	"SrcDir":          "src",
	"OutDir":          "dist",
	"AssetsDir":       "assets",
	"DevServerHost":   "localhost",
	"DevServerPort":   "3000",
}

func setConfigDefaults(cfg *ViteConfig) {
	if cfg.Env == "" {
		cfg.Env = defaults["Env"]
	}

	if cfg.Platform == "" {
		cfg.Platform = defaults["Platform"]
	}

	if cfg.MainEntry == "" {
		cfg.MainEntry = defaults["MainEntry"]
	}

	if cfg.AssetsURLPrefix == "" {
		cfg.AssetsURLPrefix = defaults["AssetsURLPrefix"]
	}

	if cfg.AssetsDir == "" {
		cfg.AssetsDir = defaults["AssetsDir"]
	}

	if cfg.Env == "development" {
		if cfg.DevServerHost == "" {
			cfg.DevServerHost = defaults["DevServerHost"]
		}

		if cfg.DevServerPort == "" {
			cfg.DevServerPort = defaults["DevServerPort"]
		}

		if cfg.SrcDir == "" {
			cfg.SrcDir = defaults["SrcDir"]
		}

	}
}
