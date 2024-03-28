package internal

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/lucasodra/chatstodo-client/internal/backend"
	"github.com/lucasodra/chatstodo-client/internal/constants"
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

// /htmx/reload
func (server *Server) htmxReloadData(writer http.ResponseWriter,
	request *http.Request) {
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)
	log.Println(backend.RefreshData(session))

	writer.Header().Set("Hx-Trigger", "reloadData")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(""))
}
