package internal

import (
	"github.com/gorilla/csrf"
	"html/template"
	"log"
	"net/http"
)

func (server *Server) dashboardHome(writer http.ResponseWriter,
	request *http.Request) {
	dashboardHandler(writer, TabListEntry{
		Id:          "tab-home",
		Title:       "home",
		RedirectUrl: "/home",
		HtmxPath:    "/htmx/home",
	})
}

// Dashboard home panel, lazy loading
func (server *Server) htmxHomePanel(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/home.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
	})
	log.Println("served home")
}
