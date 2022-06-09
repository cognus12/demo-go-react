package web

import (
	"html/template"
)

func LoadTemplate() *template.Template {
	t, err := template.ParseFiles("template.html")

	if err == nil {
		return t
	} else {
		panic(err)
	}
}
