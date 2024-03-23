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
	// Error ignored due to auth wrapper taking care of it
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)
	tasks, events, summaries := backend.GetSummary(writer, request, session)

	tmpl, err := template.ParseFiles("./templates/htmx/home.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	taskTags, eventTags, summaryTags := extractTags(tasks, events, summaries)

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
		"tasks":          tasks,
		"taskTags":       taskTags,

		"events":    events,
		"eventTags": eventTags,

		"summaries":   summaries,
		"summaryTags": summaryTags,
	})
	log.Println("served home")
}

func extractTags(tasks []*backend.Task, events []*backend.Event,
	summaries []*backend.Summary) ([]string, []string, []string) {

	tagsMap := map[string]struct{}{}
	for _, task := range tasks {
		for _, tag := range task.Tags {
			tagsMap[tag] = struct{}{}
		}
	}
	taskTagsArray := []string{}
	for key := range tagsMap {
		taskTagsArray = append(taskTagsArray, key)
	}

	tagsMap = map[string]struct{}{}
	for _, event := range events {
		for _, tag := range event.Tags {
			tagsMap[tag] = struct{}{}
		}
	}
	eventTagsArray := []string{}
	for key := range tagsMap {
		eventTagsArray = append(eventTagsArray, key)
	}

	tagsMap = map[string]struct{}{}
	for _, summary := range summaries {
		for _, tag := range summary.Tags {
			tagsMap[tag] = struct{}{}
		}
	}
	summaryTagsArray := []string{}
	for key := range tagsMap {
		summaryTagsArray = append(summaryTagsArray, key)
	}

	return taskTagsArray, eventTagsArray, summaryTagsArray
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
