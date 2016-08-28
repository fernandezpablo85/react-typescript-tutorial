package main

import (
	"html/template"
	"net/http"
	"strings"
)

//go:generate go-bindata views/...
var templates = template.New("")

func init() {
	for _, name := range AssetNames() {
		bytes, _ := Asset(name)
		logger.Debugf("registering template: %s", templateName(name))
		templates.New(templateName(name)).Parse(string(bytes))
	}
}

func templateName(path string) string {
	dot := strings.LastIndex(path, ".")
	slash := strings.LastIndex(path, "/")
	return path[slash+1 : dot]
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		logger.WithField("template", name).WithField("error", err).Warnf("error while rendering template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
