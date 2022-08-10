package main

import (
	"demo-go-react/internal/hello"
	"demo-go-react/pkg/vite"
	"html/template"
	"log"
	"net/http"
	"os"
)

var index *template.Template

// for production mode remove space before go:embed, uncomment var
// go:embed frontend/dist template.html
// var frontend embed.FS

// for dev mode
var frontend = os.DirFS("frontend")

var Config = vite.ViteConfig{
	FS:         frontend,
	ProjectDir: "frontend",
	OutDir:     "dist",
	// Env:        "development",
}

func main() {
	log.Println("Start server on localhost:8000")

	// any custom variables to be passed to template
	data := map[string]any{
		"title": "Go-React App",
	}

	v, err := vite.NewVite(&Config, data)

	if err != nil {
		log.Fatal(err)
	}

	var templateErr error

	// for prod
	// index, templateErr = template.ParseFS(frontend, "template.html")

	index, templateErr = template.ParseFiles("template.html")

	if templateErr != nil {
		log.Fatal("Template loading error: ", templateErr)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			log.Println("Execute index.html")
			err := index.Execute(w, v)

			if err != nil {
				log.Fatal("Template exicuting error:", err)
			}
		}
	})

	// set assets handler
	http.Handle(v.AssetsURLPrefix, v.FileServer())

	// handle demo rest endpoit
	http.HandleFunc("/api/hello", hello.SayHello)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
