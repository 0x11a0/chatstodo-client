package internal

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/csrf"
)

/*
type Event struct {
	Id        int      `json:"id"`
	Value     string   `json:"value"`
	Location  string   `json:"location"`
	DateStart string   `json:"dateStart"`
	DateEnd   string   `json:"dateEnd"`
	Tags      []string `json:"tags"`
}
*/

type Event1 struct {
	Id          int
	Title       string
	Date        string
	DateTime    time.Time
	Description string
}

var fakeHomeEventsData = []Event1{
	{
		Id:          0,
		Title:       "IDP outing at KSTAR",
		Date:        "17/03/2024",
		Description: "Description1",
	},
	{
		Id:          1,
		Title:       "Antoine outing at LAVO",
		Date:        "18/03/2024",
		Description: "Description2",
	},
}

// /htmx/home/events
func (server *Server) htmxEvents(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		htmxEventCard(writer, request)
	} else if request.Method == "POST" {
		htmxEventModal(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

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
	title := request.FormValue("title")
	date := request.FormValue("date")
	description := request.FormValue("description")

	//log.Println(id, title, date, description)

	tmpl, err := template.ParseFiles("./templates/htmx/eventModal.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, Todo1{
		Id:          id,
		Title:       title,
		Date:        date,
		Description: description,
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
	title := request.FormValue("title")
	date := request.FormValue("date")
	description := request.FormValue("description")
	for _, v := range fakeHomeEventsData {
		if v.Id == id {
			v.Title = title
			v.Date = date
			v.Description = description
			break
		}
	}

	tmpl, err := template.ParseFiles("./templates/htmx/eventTemplate.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, Todo1{
		Id:          id,
		Title:       title,
		Date:        date,
		Description: description,
	})
}
