package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func handleMainPage(w http.ResponseWriter, _ *http.Request) {
	var data interface{}
	_ = t.ExecuteTemplate(w, "index.html.tmpl", data)
}
