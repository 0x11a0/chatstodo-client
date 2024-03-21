package internal

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/lucasodra/chatstodo-client/internal/backend"
	"github.com/lucasodra/chatstodo-client/internal/utils"
)

// /htmx/home/tasks
func (server *Server) htmxTasks(writer http.ResponseWriter,
	request *http.Request) {
	/*
		if request.Method == "GET" {
			htmxEventCard(writer, request)
		} else
	*/
	if request.Method == "POST" {
		htmxTaskModal(writer, request)
	} else if request.Method == "PUT" {
		htmxTaskSave(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
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

/*


   <div class="card-date-container">
       <label>Deadline:</label>
       <input type="text" name="deadline" value="{{ .DisplayDeadline }}" class="card-block-date"
           readonly required></input>
   </div>

   <input type="text" name="tags" value="{{ .DisplayTags }}" class="card-block-notes" readonly
       required></input>

   <input type="hidden" name="localDeadline" value="{{ .HTMLDeadline }}" required></input>
*/

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
