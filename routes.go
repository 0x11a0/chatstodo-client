package main

import (
	"html/template"
	"net/http"
	_ "time"
)

func indexHandler(writer http.ResponseWriter, request *http.Request,
	htmxPath string, redirectUrl string) {
	tmpl, err := template.ParseFiles("./templates/index.html", "./templates/dashboard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(writer, "base", map[string]interface{}{
		"htmxPath":    htmxPath,
		"redirectUrl": redirectUrl,
	})
}

func (server *Server) indexHome(writer http.ResponseWriter,
	request *http.Request) {
	indexHandler(writer, request, "/htmx/home", "/home")
}

func (server *Server) indexBots(writer http.ResponseWriter,
	request *http.Request) {
	indexHandler(writer, request, "/htmx/bots", "/bots")
}

func (server *Server) indexSettings(writer http.ResponseWriter,
	request *http.Request) {
	indexHandler(writer, request, "/htmx/settings", "/settings")
}

func (server *Server) htmxHomePanel(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/htmx/home.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, nil)
}

func (server *Server) loginPage(writer http.ResponseWriter,
	request *http.Request) {
	if server.isValidSession(request) {
		http.Redirect(writer, request, "/home", http.StatusSeeOther)
		//server.indexHome(writer, request)
		return
	}
	tmpl, err := template.ParseFiles("./templates/index.html", "./templates/login.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(writer, "base", nil)
}

func (server *Server) htmxSettings(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/settings.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{})
}
