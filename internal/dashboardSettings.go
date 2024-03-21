package internal

import (
	"html/template"
	"net/http"
)

func (server *Server) dashboardSettings(writer http.ResponseWriter,
	request *http.Request) {
	dashboardHandler(writer, "/htmx/settings", "/settings")
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
