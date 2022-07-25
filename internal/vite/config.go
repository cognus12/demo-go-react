package vite

type ViteConfig struct {
	// relative path, eg static/assets
	AssetsPath string
	// relative path, eg static/manifest.json
	ManifestPath string
	// production or development, default - production
	Env string
	// react, vue, svelte, default - react
	Platform string
	// main chunk name, default - src/main.tsx
	MainEntry string
}

var defaults = map[string]string{
	"Env":       "production",
	"Platform":  "react",
	"MainEntry": "src/main.tsx",
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
}
