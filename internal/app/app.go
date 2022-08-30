package app

import (
	"demo-go-react/internal/hello"
	"demo-go-react/pkg/vite"
	"log"
	"net/http"
)

func Run(cfg *vite.ViteConfig) {
	v, err := vite.NewVite(cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start server on localhost:8000")

	data := map[string]any{
		"title": "Go-React App",
	}

	// any custom variables to be passed to template
	err = v.SetArgs(data)

	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			log.Println("Execute index.html")
			err := v.Template.Execute(w, v)

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
