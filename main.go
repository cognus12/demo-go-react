package main

import (
	"demo-go-react/internal/app"
	"demo-go-react/pkg/vite"
	"embed"
	"html/template"
	"log"
)

var index *template.Template

// for development need to remove /dist (if there is no dist folder at the moment)
//go:embed frontend/dist template.html
var frontend embed.FS

var Config = vite.ViteConfig{
	FS:         frontend,
	ProjectDir: "frontend",
	OutDir:     "dist",
	Env:        "development",
}

func main() {
	var templateErr error
	Config.Template, templateErr = template.ParseFS(frontend, "template.html")

	if templateErr != nil {
		log.Fatal("Template loading error: ", templateErr)
	}

	app.Run(&Config)
}
