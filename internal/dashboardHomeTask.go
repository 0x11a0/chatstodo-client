package internal

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/lucasodra/chatstodo-client/internal/backend"
	"github.com/lucasodra/chatstodo-client/internal/constants"
	"github.com/lucasodra/chatstodo-client/internal/utils"
	"google.golang.org/api/calendar/v3"
)

// /htmx/home/tasks/export
func (server *Server) htmxTasksExport(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		//server.htmxHomeTasks(writer, request)
	} else if request.Method == "POST" {
		htmxTaskExportModal(writer, request)
	} else if request.Method == "PUT" {
		server.htmxTaskExport(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

// /htmx/home/tasks/export "POST"
// Spawns the modal for exporting
func htmxTaskExportModal(writer http.ResponseWriter,
	request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println("dashboardHomeTask.go - htmxTaskModal(), parse form")
		log.Println(err)
		return
	}

	type TaskEntry struct {
		Value    string
		Deadline string
		Pretty   string
	}

	taskMap := map[string]*TaskEntry{}
	for key := range request.Form {
		if strings.HasPrefix(key, "value-") {
			taskId, _ := strings.CutPrefix(key, "value-")
			if taskMap[taskId] == nil {
				taskMap[taskId] = &TaskEntry{}
			}
			taskMap[taskId].Value = request.FormValue(key)
		} else if strings.HasPrefix(key, "deadline-") {
			taskId, _ := strings.CutPrefix(key, "deadline-")
			if taskMap[taskId] == nil {
				taskMap[taskId] = &TaskEntry{}
			}
			taskMap[taskId].Deadline = request.FormValue(key)
		} else if strings.HasPrefix(key, "pretty-") {
			taskId, _ := strings.CutPrefix(key, "pretty-")
			if taskMap[taskId] == nil {
				taskMap[taskId] = &TaskEntry{}
			}
			taskMap[taskId].Pretty = request.FormValue(key)
		}
	}
	tmpl, err := template.ParseFiles("./templates/htmx/taskExportModal.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"csrfToken":      csrf.Token(request),
		csrf.TemplateTag: csrf.TemplateField(request),
		"taskMap":        taskMap,
	})
}

// /htmx/home/tasks/export "PUT"
// Exports to external calendar.
// WARN: Currently hardcoded for Google
// calendar and for Singapore timezone
func (server *Server) htmxTaskExport(writer http.ResponseWriter,
	request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println("dashboardHomeTask.go - htmxTaskExport(), parse form")
		log.Println(err)
		return
	}
	type TaskEntry struct {
		Value    string `json:"value"`
		Deadline string `json:"deadline"`
	}
	var tasks []*TaskEntry
	err = json.NewDecoder(request.Body).Decode(&tasks)
	if err != nil {
		log.Println("dashboardHomeTask.go - htmxTaskExport(), decode")
		log.Println(err)
		return

	}

	for _, task := range tasks {
		deadlineStr := utils.HTMLToGCalendar(task.Deadline)
		calendarEvent := &calendar.Event{
			Summary:     task.Value,
			Location:    "",
			Description: "",
			Start: &calendar.EventDateTime{
				DateTime: deadlineStr,
				TimeZone: "Asia/Singapore",
			},
			End: &calendar.EventDateTime{
				DateTime: deadlineStr,
				TimeZone: "Asia/Singapore",
			},
			Recurrence: []string{},
			Attendees:  []*calendar.EventAttendee{},
		}
		server.exportEvent(writer, request, calendarEvent)
	}
}

// /htmx/home/tasks
func (server *Server) htmxTasks(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		server.htmxHomeTasks(writer, request)
	} else if request.Method == "POST" {
		htmxTaskModal(writer, request)
	} else if request.Method == "PUT" {
		htmxTaskSave(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

// /htmx/home/tasks "GET"
func (server *Server) htmxHomeTasks(writer http.ResponseWriter,
	request *http.Request) {
	// Error ignored due to auth wrapper taking care of it
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)
	// TODO: Error handling
	tasks, _ := backend.GetTasks(writer, request, session)

	tagsMap := map[string]struct{}{}
	for _, task := range tasks {
		for _, tag := range task.Tags {
			tagsMap[tag] = struct{}{}
		}
	}
	taskTags := []string{}
	for key := range tagsMap {
		taskTags = append(taskTags, key)
	}
	slices.Sort(taskTags)

	tmpl, err := template.ParseFiles("./templates/htmx/taskCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		"csrfToken":      csrf.Token(request),
		csrf.TemplateTag: csrf.TemplateField(request),
		"tasks":          tasks,
		"taskTags":       taskTags,
	})

}

// /htmx/home/tasks "POST"
// Spawns the modal for editing of task details
func htmxTaskModal(writer http.ResponseWriter,
	request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println("dashboardHomeTask.go - htmxTaskModal(), parse form")
		log.Println(err)
		return
	}
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "Mismatched data found in form.", http.StatusUnauthorized)
		return
	}
	value := request.FormValue("value")
	deadline := request.FormValue("localDeadline")
	tags := request.FormValue("tags")
	tmpl, err := template.ParseFiles("./templates/htmx/taskEditModal.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
		"Id":             id,
		"Value":          value,
		"HTMLDeadline":   deadline,
		"DisplayTags": strings.TrimSuffix(
			strings.TrimPrefix(tags, "["),
			"]"),
	})
}

// /htmx/home/tasks "PUT"
// Saves the edited information into this
func htmxTaskSave(writer http.ResponseWriter,
	request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println("dashboardHomeTask.go - htmxTaskSave()")
		log.Println(err)
		return
	}
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "Mismatched data found in form.", http.StatusUnauthorized)
		return
	}

	task := backend.Task{
		Id:           id,
		Value:        request.FormValue("value"),
		HTMLDeadline: request.FormValue("deadline"),
	}
	task.DisplayDeadline = utils.PrettifyHTMLDateTime(task.HTMLDeadline)
	newTags := strings.Split(request.FormValue("tags"), ",")
	for i := 0; i < len(newTags); i++ {
		newTags[i] = strings.TrimSpace(newTags[i])
	}
	log.Println(newTags)
	task.Tags = newTags
	task.DisplayTags = "[" + strings.Join(newTags, ", ") + "]"

	// TODO: Update backend about latest data

	tmpl, err := template.ParseFiles("./templates/htmx/taskTemplate.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag:  csrf.TemplateField(request),
		"Id":              task.Id,
		"Value":           task.Value,
		"Tags":            task.Tags,
		"DisplayTags":     task.DisplayTags,
		"HTMLDeadline":    task.HTMLDeadline,
		"DisplayDeadline": task.DisplayDeadline,
	})
}
