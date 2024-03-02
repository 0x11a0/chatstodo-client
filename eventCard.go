package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/markbates/goth/gothic"
)

type Event struct {
	Id          int
	Title       string
	Date        string
	Description string
}

var fakeHomeEventsData = []Event{
	{
		Id:          0,
		Title:       "IDP outing at KSTAR",
		Date:        "17/03/2024",
		Description: "Description",
	},
	{
		Id:          1,
		Title:       "Antoine outing at LAVO",
		Date:        "18/03/2024",
		Description: "Description",
	},
}

func (server *Server) htmxEventCard(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		htmxEventCardAll(writer, request)
	} else if request.Method == "POST" {
		htmxEventSave(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func htmxEventCardAll(writer http.ResponseWriter,
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

func htmxEventSave(writer http.ResponseWriter,
	request *http.Request) {
	_, err := gothic.GetFromSession("userId", request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	err = request.ParseForm()
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
	tmpl.Execute(writer, Todo{
		Id:          id,
		Title:       title,
		Date:        date,
		Description: description,
	})
}
