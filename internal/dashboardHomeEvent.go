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

// /htmx/home/events
func (server *Server) htmxEvents(writer http.ResponseWriter,
	request *http.Request) {
	/*
		if request.Method == "GET" {
			htmxEventCard(writer, request)
		} else
	*/
	if request.Method == "POST" {
		htmxEventModal(writer, request)
	} else if request.Method == "PUT" {
		htmxEventSave(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

/*
// /htmx/home/events "GET"
func htmxEventCard(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/eventCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"events":         fakeHomeEventsData,
		csrf.TemplateTag: csrf.TemplateField(request),
	})
}
*/

// /htmx/home/events "POST"
// Spawns the modal for editing of event details
func htmxEventModal(writer http.ResponseWriter,
	request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println("todoCard.go - htmxEventModal(), parse form")
		log.Println(err)
		return
	}
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "Mismatched data found in form.", http.StatusUnauthorized)
		return
	}
	value := request.FormValue("value")
	eventLocation := request.FormValue("location")
	dateStart := request.FormValue("localDateStart")
	dateEnd := request.FormValue("localDateEnd")
	tags := request.FormValue("tags")
	tmpl, err := template.ParseFiles("./templates/htmx/eventModal.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
		"Id":             id,
		"Value":          value,
		"Location":       eventLocation,
		"HTMLDateStart":  dateStart,
		"HTMLDateEnd":    dateEnd,
		"DisplayTags": strings.TrimSuffix(
			strings.TrimPrefix(tags, "["),
			"]"),
	})
}

// /htmx/home/events "PUT"
// Saves the edited information into this
func htmxEventSave(writer http.ResponseWriter,
	request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println("todoCard.go - htmxEventSave()")
		log.Println(err)
		return
	}
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "Mismatched data found in form.", http.StatusUnauthorized)
		return
	}

	/*
			   <input name="id" type="hidden" value="{{ .Id }}"></input>

			   <div class="mui-textfield mui-textfield--float-label">
			       <input name="value" type="text" value="{{ .Value }}"></input>
			       <label>Title</label>
			   </div>
			   <div class="mui-textfield mui-textfield--float-label">
			       <input name="location" type="text" value="{{ .Location }}"></input>
			       <label>Location</label>
			   </div>
			   <div class="mui-textfield mui-textfield--float-label">
			       <input id="modal-event-date-start" name="dateStart" type="datetime-local"
			           value="{{ .LocalDateStart }}"></input>
			       <label>Start date</label>
			   </div>
			   <div class="mui-textfield mui-textfield--float-label">
			       <input id="modal-event-date-end" name="dateEnd" type="datetime-local"
			           value="{{ .LocalDateEnd }}"></input>
			       <label>End date</label>
			   </div>





		                        <input type="text" name="value" value="{{ .Value }}" class="card-block-point" readonly></input>
		                        <label>Start date:</label>
		                        <input type="text" name="displayDateStart" value="{{ .DisplayDateStart }}"
		                            class="card-block-date" readonly></input>
		                        <label>End date:</label>
		                        <input type="text" name="displayDateEnd" value="{{ .DisplayDateEnd }}" class="card-block-date"
		                            readonly></input>
		                        <input type="text" name="location" value="{{ .Location }}" class="card-block-notes"
		                            readonly></input>
		                        <p>{{ .DisplayTags }}</p>
		                        <input type="hidden" name="localDateStart" value="{{ .LocalDateStart }}"></input>
		                        <input type="hidden" name="localDateEnd" value="{{ .LocalDateEnd }}"></input>
	*/
	event := backend.Event{
		Id:            id,
		Value:         request.FormValue("value"),
		Location:      request.FormValue("location"),
		HTMLDateStart: request.FormValue("dateStart"),
		HTMLDateEnd:   request.FormValue("dateEnd"),
	}
	event.DisplayDateStart = utils.PrettifyHTMLDateTime(event.HTMLDateStart)
	event.DisplayDateEnd = utils.PrettifyHTMLDateTime(event.HTMLDateEnd)
	newTags := strings.Split(request.FormValue("tags"), ",")
	for i := 0; i < len(newTags); i++ {
		newTags[i] = strings.TrimSpace(newTags[i])
	}
	log.Println(newTags)
	event.Tags = newTags
	event.DisplayTags = "[" + strings.Join(newTags, ", ") + "]"

	// TODO: Update backend about latest data

	tmpl, err := template.ParseFiles("./templates/htmx/eventTemplate.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag:   csrf.TemplateField(request),
		"Id":               event.Id,
		"Value":            event.Value,
		"Location":         event.Location,
		"Tags":             event.Tags,
		"DisplayTags":      event.DisplayTags,
		"HTMLDateStart":    event.HTMLDateStart,
		"DisplayDateStart": event.DisplayDateStart,
		"HTMLDateEnd":      event.HTMLDateEnd,
		"DisplayDateEnd":   event.DisplayDateEnd,
	})
}