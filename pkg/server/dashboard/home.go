package dashboard

import (
	"html/template"
	"net/http"
)

func (server *Server) dashboardHome(writer http.ResponseWriter,
	request *http.Request) {
	dashboardHandler(writer, "/htmx/home", "/home")
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

// /htmx/bots/modal
func (server *Server) htmxHomeModal(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/htmx/eventModal.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{})
}
