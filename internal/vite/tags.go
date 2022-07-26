package vite

import (
	"bytes"
	"html/template"
	"strings"
)

func (v *Vite) Tags() (template.HTML, error) {
	b := strings.Builder{}

	if v.Env == "production" {
		b.WriteString(`
			<script type="module" crossorigin src="/{{ .MainEntry }}"></script>
			
			{{ range $href := .Imports }}
				<link rel="modulepreload" href="/{{ $href }}">
			{{ end }}
			
			{{ range $href := .CSS }}
				<link rel="stylesheet" href="/{{ $href }}">
			{{ end }}
		`)
	}

	tmpl, err := template.New("tags").Parse(b.String())

	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	tmpl.Execute(&buffer, v)

	return template.HTML(buffer.String()), nil
}
