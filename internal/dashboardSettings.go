package internal

import (
	"html/template"
	"log"
	"net/http"

	"github.com/lucasodra/chatstodo-client/internal/backend"
	"github.com/lucasodra/chatstodo-client/internal/constants"
)

// /settings (index handler)
func (server *Server) dashboardSettings(writer http.ResponseWriter,
	request *http.Request) {
	dashboardHandler(writer, "/htmx/settings", "/settings")
}

// /htmx/settings
func (server *Server) htmxSettings(writer http.ResponseWriter,
	request *http.Request) {

	// Will not error here due to auth wrapper
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)

	platformEntries, statusCode := backend.GetAllPlatforms(session)
	if statusCode != http.StatusOK {
		log.Println("dashboardSettings.go - htmxSettings(), get all platforms")
		return
	}

	tmpl, err := template.ParseFiles("./templates/htmx/settings.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		"platformEntries": platformEntries,
	})
}
