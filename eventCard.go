package main

import (
	"html/template"
	"net/http"
)

type Event struct {
	Title       string
	Date        string
	Description string
}

var fakeHomeEventsData = []Event{
	{
		Title:       "IDP outing at KSTAR",
		Date:        "17/03/2024",
		Description: "Description",
	},
	{
		Title:       "Antoine outing at LAVO",
		Date:        "18/03/2024",
		Description: "Description",
	},
}

func (server *Server) htmxEventCard(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/htmx/eventCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"events": fakeHomeEventsData,
	})
}
