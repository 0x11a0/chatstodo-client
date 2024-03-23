package internal

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/lucasodra/chatstodo-client/internal/backend"
	"github.com/lucasodra/chatstodo-client/internal/constants"
)

/*
	type BotEntry struct {
		Id       int
		Name     string
		Platform string
		Status   string
	}

	var fakeBotData = []BotEntry{
		{
			Id:       0,
			Name:     "Product Devvo",
			Platform: "Telegram",
			Status:   "Active",
		},
		{
			Id:       1,
			Name:     "Scrum Masters",
			Platform: "Discord",
			Status:   "Inactive",
		},
	}
*/
func (server *Server) dashboardGroups(writer http.ResponseWriter,
	request *http.Request) {
	dashboardHandler(writer, TabListEntry{
		Id:          "tab-groups",
		Title:       "Groups",
		RedirectUrl: "/groups",
		HtmxPath:    "/htmx/groups",
	})
}

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
