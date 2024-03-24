package internal

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/csrf"
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

var (
	SUPPORTED_PLATFORMS = map[string]bool{
		"Telegram": false,
		"Discord":  false,
	}
)

// /htmx/settings "GET"
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

// /htmx/settings/platform
func (server *Server) htmxSettingsPlatform(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		server.htmxSettingsCodeModal(writer, request)
	} else if request.Method == "POST" {
		server.htmxSettingsAddPlatform(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

// /htmx/settings/platform "GET"
func (server *Server) htmxSettingsCodeModal(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/settingsCode.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
		".inputValue":    "",
	})

}

var (
	ADD_PLATFORM_ERROR_MESSAGES = map[int]string{
		http.StatusOK:                  "Your account has been successfully linked!",
		http.StatusNotFound:            "Invalid verification code, or code may have expired. Please request again!",
		http.StatusConflict:            "You have already linked this platform on this account!",
		http.StatusInternalServerError: "An error occured with our servers, please try again later!",
	}
)

// /htmx/settings/platform "POST"
func (server *Server) htmxSettingsAddPlatform(writer http.ResponseWriter,
	request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println("dashboardSettings.go - htmxSettingsAddPlatform(), parse form")
		log.Println(err)
		return
	}
	verificationCode := strings.TrimSpace(request.FormValue("verificationCode"))
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)
	statusCode := backend.AddPlatform(session, verificationCode)
	log.Println("addPlatform statusCode", statusCode)
	// TODO:
	if statusCode == http.StatusUnauthorized {

	}

	tmpl, err := template.ParseFiles("./templates/htmx/settingsCode.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
		"errorMessage":   ADD_PLATFORM_ERROR_MESSAGES[statusCode],
	})

}
