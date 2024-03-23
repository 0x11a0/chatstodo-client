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
	dashboardHandler(writer, TabListEntry{
		Id:          "tab-settings",
		Title:       "Settings",
		RedirectUrl: "/settings",
		HtmxPath:    "/htmx/settings",
	})
}

var (
	SUPPORTED_PLATFORMS = map[string]bool{
		"Telegram": false,
		"Discord":  false,
	}
)

/*

// /htmx/groups for both tab panel and search functionality
func (server *Server) htmxGroups(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		server.htmxGroupsPanel(writer, request)
	} else if request.Method == "POST" {
		//htmxBotsSearch(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

// /htmx/groups "GET" for tab panel
func (server *Server) htmxGroupsPanel(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/groups.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)
	groups, statusCode := backend.GetAllGroups(*session)
	if statusCode != http.StatusOK {
		log.Println("dashboardGroups.go htmxGroupsPanel(), backend error status ")
		log.Println(statusCode)
		return
	}
	log.Println(groups)

	tmpl.Execute(writer, map[string]interface{}{
		"csrfToken": csrf.Token(request),
	})
}

*/

// /htmx/settings
func (server *Server) htmxSettings(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		server.htmxSettingsPanel(writer, request)
	} else if request.Method == "POST" {
		//htmxBotsSearch(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func (server *Server) htmxSettingsPanel(writer http.ResponseWriter,
	request *http.Request) {

	// Will not error here due to auth wrapper
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)

	platformEntries, statusCode := backend.GetAllPlatforms(session)
	if statusCode != http.StatusOK {
		log.Println("dashboardSettings.go - htmxSettings(), get all platforms")
		return
	}

	platformMap := map[string]bool{}
	for key, value := range SUPPORTED_PLATFORMS {
		platformMap[key] = value
	}

	for _, platform := range platformEntries {
		platformMap[platform.PlatformName] = true
	}

	tmpl, err := template.ParseFiles("./templates/htmx/settings.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		"platforms": platformMap,
	})

}
