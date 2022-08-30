package main

import (
	"demo-go-react/internal/app"
	"demo-go-react/pkg/vite"
	"embed"
	"html/template"
	"log"
)

var index *template.Template

//go:embed static template.html
var fs embed.FS

var Config = vite.ViteConfig{
	RootFS: fs,
	// Env: "development",
}

func main() {
	var templateErr error
	Config.Template, templateErr = template.ParseFS(fs, "template.html")

	if templateErr != nil {
		log.Fatal("Template loading error: ", templateErr)
	}

	app.Run(&Config)
}
