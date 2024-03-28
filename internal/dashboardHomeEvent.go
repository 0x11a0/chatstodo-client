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

// /htmx/home/events/export
func (server *Server) htmxEventsExport(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		//server.htmxHomeTasks(writer, request)
	} else if request.Method == "POST" {
		htmxEventsExportModal(writer, request)
	} else if request.Method == "PUT" {
		server.htmxEventExport(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

var (
	EVENTS_EXPORT_PREFIXES = []string{
		"value-",
		"location-",
		"startDate-",
		"prettyStart-",
		"endDate-",
		"prettyEnd-",
	}
)

// /htmx/home/events/export "POST"
// Spawns the modal for exporting
func htmxEventsExportModal(writer http.ResponseWriter,
	request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println("dashboardHomeEvent.go - htmxEventsExportModal(), parse form")
		log.Println(err)
		return
	}

	type EventEntry struct {
		Value       string
		Location    string
		StartDate   string
		PrettyStart string
		EndDate     string
		PrettyEnd   string
	}

	eventMap := map[string]*EventEntry{}
	for key := range request.Form {
		for i, prefix := range EVENTS_EXPORT_PREFIXES {
			if strings.HasPrefix(key, prefix) {
				eventId, _ := strings.CutPrefix(key, prefix)
				if eventMap[eventId] == nil {
					eventMap[eventId] = &EventEntry{}
				}
				switch i {
				case 0:
					eventMap[eventId].Value = request.FormValue(key)
				case 1:
					eventMap[eventId].Location = request.FormValue(key)
				case 2:
					eventMap[eventId].StartDate = request.FormValue(key)
				case 3:
					eventMap[eventId].PrettyStart = request.FormValue(key)
				case 4:
					eventMap[eventId].EndDate = request.FormValue(key)
				case 5:
					eventMap[eventId].PrettyEnd = request.FormValue(key)
				}
				break
			}
		}
	}

	tmpl, err := template.ParseFiles("./templates/htmx/eventExportModal.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"csrfToken":      csrf.Token(request),
		csrf.TemplateTag: csrf.TemplateField(request),
		"eventMap":       eventMap,
	})
}

// /htmx/home/events/export "PUT"
// Exports to external calendar.
// WARN: Currently hardcoded for Google
// calendar and for Singapore timezone
func (server *Server) htmxEventExport(writer http.ResponseWriter,
	request *http.Request) {

	// NOTE: Might be redundant
	err := request.ParseForm()
	if err != nil {
		log.Println("dashboardHomeEvent.go - htmxEventExport(), parse form")
		log.Println(err)
		return
	}

	type EventEntry struct {
		Value     string `json:"value"`
		Location  string `json:"location"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}
	var events []*EventEntry
	err = json.NewDecoder(request.Body).Decode(&events)
	if err != nil {
		log.Println("dashboardHomeEvent.go - htmxEventExport(), decode")
		log.Println(err)
		return
	}

	for _, event := range events {
		startStr := utils.HTMLToGCalendar(event.StartDate)
		endStr := utils.HTMLToGCalendar(event.EndDate)
		calendarEvent := &calendar.Event{
			Summary:     event.Value,
			Location:    event.Location,
			Description: "",
			Start: &calendar.EventDateTime{
				DateTime: startStr,
				TimeZone: "Asia/Singapore",
			},
			End: &calendar.EventDateTime{
				DateTime: endStr,
				TimeZone: "Asia/Singapore",
			},
			Recurrence: []string{},
			Attendees:  []*calendar.EventAttendee{},
		}
		server.exportEvent(writer, request, calendarEvent)
	}
}

// /htmx/home/events
func (server *Server) htmxEvents(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		server.htmxEventCard(writer, request)
	} else if request.Method == "POST" {
		htmxEventModal(writer, request)
	} else if request.Method == "PUT" {
		htmxEventSave(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

// /htmx/home/events "GET"
func (server *Server) htmxEventCard(writer http.ResponseWriter,
	request *http.Request) {
	// Error ignored due to auth wrapper taking care of it
	session, _ := server.redisSessionStore.Get(request, constants.COOKIE_NAME)
	// TODO: Error handling
	events, _ := backend.GetEvents(session)

	eventsMap := map[string]struct{}{}
	for _, task := range events {
		for _, tag := range task.Tags {
			eventsMap[tag] = struct{}{}
		}
	}
	eventTags := []string{}
	for key := range eventsMap {
		eventTags = append(eventTags, key)
	}
	slices.Sort(eventTags)

	tmpl, err := template.ParseFiles("./templates/htmx/eventCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
		"events":         events,
		"eventTags":      eventTags,
	})

}

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
	tmpl, err := template.ParseFiles("./templates/htmx/eventEditModal.html")
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
