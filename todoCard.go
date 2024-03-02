package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/markbates/goth/gothic"
)

type Todo struct {
	Id          int
	Title       string
	Date        string
	Description string
	Urgent      bool
}

var fakeHomeTodoData = []*Todo{
	{
		Id:          0,
		Title:       "Fill indemnity form",
		Date:        "14/02/2024",
		Description: "Needs to be completed ASAP",
		Urgent:      true,
	},
	{
		Id:     1,
		Title:  "Research places with gin tonic",
		Date:   "11/02/2024",
		Urgent: false,
	},
	{
		Id:     2,
		Title:  "Make a reservation for LAVO on March 18, 3pm",
		Date:   "11/02/2024",
		Urgent: false,
	},
	{
		Id:     3,
		Title:  "Bring cake for the celebration",
		Date:   "11/02/2024",
		Urgent: false,
	},
}

func (server *Server) htmxTodoCard(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		htmxTodoCardAll(writer, request)
	} else if request.Method == "POST" {
		htmxTodoSave(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func htmxTodoCardAll(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/todoCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"todos":          fakeHomeTodoData,
		csrf.TemplateTag: csrf.TemplateField(request),
	})
}

func htmxTodoSave(writer http.ResponseWriter,
	request *http.Request) {
	_, err := gothic.GetFromSession("userId", request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	err = request.ParseForm()
	if err != nil {
		log.Println("todoCard.go - htmxTodoSave()")
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
	for _, v := range fakeHomeTodoData {
		if v.Id == id {
			v.Title = title
			v.Date = date
			v.Description = description
			break
		}
	}

	tmpl, err := template.ParseFiles("./templates/htmx/todoTemplate.html")
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
