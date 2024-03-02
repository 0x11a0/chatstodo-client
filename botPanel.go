package main

import (
	"html/template"
	"net/http"
)

type BotEntry struct {
	BotType  string
	Name     string
	IsActive bool
}

var fakeBotData = []BotEntry{
	{
		BotType:  "telegram",
		Name:     "Telegram Chat 1",
		IsActive: true,
	},
	{
		BotType:  "telegram",
		Name:     "Telegram Chat 2",
		IsActive: false,
	},
}

func (server *Server) htmxBots(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/htmx/bots.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"bots": fakeBotData,
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
