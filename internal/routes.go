package internal

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func dashboardHandler(writer http.ResponseWriter,
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

func (server *Server) dashboardSettings(writer http.ResponseWriter,
	request *http.Request) {
	dashboardHandler(writer, "/htmx/settings", "/settings")
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

var (
	ERROR_MAP = map[int]string{

		http.StatusUnauthorized: "Unauthorized!",              // 401
		http.StatusNotFound:     "Page or content not found!", // 404

	}
)

func (server *Server) errorPageNotGeneric(writer http.ResponseWriter,
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
