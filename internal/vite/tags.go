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

	if v.Env == "development" {
		if v.Platform == "react" {
			b.WriteString(`
				<script type="module">
					import RefreshRuntime from '{{ .DevServerURL }}/@react-refresh'
					RefreshRuntime.injectIntoGlobalHook(window)
					window.$RefreshReg$ = () => {}
					window.$RefreshSig$ = () => (type) => type
					window.__vite_plugin_react_preamble_installed__ = true
				</script>
			`)
		}

		b.WriteString(`
			<script type="module" src="{{ .DevServerURL }}/@vite/client"></script>
        	<script type="module" src="{{ .DevServerURL }}/{{ .SrcDir }}/main.tsx"></script>
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
