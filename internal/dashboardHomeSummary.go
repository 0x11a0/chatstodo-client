package internal

import (
	"html/template"
	"net/http"
	"slices"
	"github.com/gorilla/csrf"
	"github.com/lucasodra/chatstodo-client/internal/backend"
	"github.com/lucasodra/chatstodo-client/internal/constants"
)

// /htmx/home/summaries
func (server *Server) htmxSummaries(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		server.htmxHomeSummaries(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

// /htmx/home/summaries "GET"
func (server *Server) htmxHomeSummaries(writer http.ResponseWriter,
	request *http.Request) {
	// Error ignored due to auth wrapper taking care of it
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)
	// TODO: Error handling
	summaries, _ := backend.GetSummaries(writer, request, session)

	summariesMap := map[string]struct{}{}
	for _, task := range summaries {
		for _, tag := range task.Tags {
			summariesMap[tag] = struct{}{}
		}
	}
	summaryTags := []string{}
	for key := range summariesMap {
		summaryTags = append(summaryTags, key)
	}
	slices.Sort(summaryTags)

	tmpl, err := template.ParseFiles("./templates/htmx/summaryCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
		"summaries":      summaries,
		"summaryTags":    summaryTags,
	})

}
