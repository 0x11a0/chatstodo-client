package internal

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TabListEntry struct {
	// html id
	Id          string
	Title       string
	HtmxPath    string
	RedirectUrl string
}

var (
	tabListEntries = []TabListEntry{
		{
			Id:          "tab-home",
			Title:       "Home",
			HtmxPath:    "/htmx/home",
			RedirectUrl: "/home",
		},
		{
			Id:          "tab-groups",
			Title:       "Groups",
			HtmxPath:    "/htmx/groups",
			RedirectUrl: "/groups",
		},
		{
			Id:          "tab-settings",
			Title:       "Settings",
			HtmxPath:    "/htmx/settings",
			RedirectUrl: "/settings",
		},
	}
)

// Renders the dashboard html once.
// E.g. for base/htmx/home, htmxPath is "/htmx/home"
func dashboardHandler(writer http.ResponseWriter,
	tabListEntry TabListEntry) {
	tmpl, err := template.ParseFiles("./templates/index.html", "./templates/dashboard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(writer, "base", map[string]interface{}{
		"tabListEntries": tabListEntries,
		"htmxPath":       tabListEntry.HtmxPath,
		"redirectUrl":    tabListEntry.RedirectUrl,
	})
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

var (
	ERROR_MAP = map[int]string{
		http.StatusUnauthorized: "Unauthorized!",              // 401
		http.StatusNotFound:     "Page or content not found!", // 404
	}
)

func (server *Server) errorPageGeneric(writer http.ResponseWriter,
	request *http.Request) {

	params := mux.Vars(request)
	statusCode, err := strconv.Atoi(params["code"])
	if err != nil {
		statusCode = http.StatusNotFound
	}
	log.Println(params)
	errorMessage := ERROR_MAP[statusCode]
	if errorMessage == "" {
		statusCode = http.StatusNotFound
	}
	errorMessage = ERROR_MAP[statusCode]

	server.errorPageSpecific(
		writer,
		statusCode,
		errorMessage,
	)
}

func (server *Server) errorPageSpecific(writer http.ResponseWriter,
	statusCode int, errorMessage string) {

	tmpl, err := template.ParseFiles("./templates/index.html", "./templates/error.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(writer, "base",
		map[string]interface{}{
			"statusCode":   statusCode,
			"errorMessage": errorMessage,
		})
}

func (server *Server) errorPageNotFound(writer http.ResponseWriter,
	_ *http.Request) {
	server.errorPageSpecific(
		writer,
		http.StatusNotFound,
		ERROR_MAP[http.StatusNotFound],
	)
}
