package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["blue-green.html"]
	homeController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
