package main

import (
	"github.com/gorilla/csrf"
	"github.com/markbates/goth/gothic"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type BotEntry struct {
	Id       int
	Name     string
	Platform string
	Status   string
}

var fakeBotData = []BotEntry{
	{
		Id:       0,
		Name:     "Product Devvo",
		Platform: "Telegram",
		Status:   "Active",
	},
	{
		Id:       1,
		Name:     "Scrum Masters",
		Platform: "Discord",
		Status:   "Inactive",
	},
}

func (server *Server) htmxBots(writer http.ResponseWriter,
	request *http.Request) {
	if request.Method == "GET" {
		htmxBotsPanel(writer, request)
	} else if request.Method == "POST" {
		htmxBotsSearch(writer, request)
	} else {
		http.Error(writer, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func htmxBotsPanel(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/bots.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"csrfToken": csrf.Token(request),
	})
}

func htmxBotsSearch(writer http.ResponseWriter,
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
	search := request.FormValue("search")
	var botData []BotEntry
	for _, bot := range fakeBotData {
		if strings.Contains(strings.ToLower(bot.Name), search) {
			botData = append(botData, bot)
		}
	}
	tmpl, err := template.ParseFiles("./templates/htmx/botSearch.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"bots": botData,
	})
}

func (server *Server) htmxBotModal(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/botModal.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{})
}
