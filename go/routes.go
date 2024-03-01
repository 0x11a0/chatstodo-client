package main

import (
	"html/template"
	"net/http"
	_ "time"
)

type User struct {
	UserId int
	Name   string
}

func indexHandler(writer http.ResponseWriter,
	htmxPath string, redirectUrl string) {

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"htmxPath":    htmxPath,
		"redirectUrl": redirectUrl,
	})
}

func (server *Server) indexHome(writer http.ResponseWriter,
	request *http.Request) {
	indexHandler(writer, "/htmx/home", "/home")
}

func (server *Server) indexBots(writer http.ResponseWriter,
	request *http.Request) {
	indexHandler(writer, "/htmx/bots", "/bots")
}
func (server *Server) indexSettings(writer http.ResponseWriter,
	request *http.Request) {
	indexHandler(writer, "/htmx/settings", "/settings")
}

type Entry struct {
	Title string
	// Date        time.Time
	Date        string
	Description string
	Urgent      bool
}

func (server *Server) homeHandler(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, nil)

}

var fakeHomeToDoData = []Entry{
	{
		Title:       "Fill indemnity form",
		Date:        "14/02/2024",
		Description: "Needs to be completed ASAP",
		Urgent:      true,
	},
	{
		Title:  "Research places with gin tonic",
		Date:   "11/02/2024",
		Urgent: false,
	},
	{
		Title:  "Make a reservation for LAVO on March 18, 3pm",
		Date:   "11/02/2024",
		Urgent: false,
	},
	{
		Title:  "Bring cake for the celebration",
		Date:   "11/02/2024",
		Urgent: false,
	},
}
var fakeHomeEventsData = []Entry{
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

type Summary struct {
	Title           string
	Date            string
	Points          []string
	ImportantPoints []string
}

var fakeHomeSummaryData = []Summary{
	{
		Title: "Validation Group 1",
		Date:  "11/02/2024",
		Points: []string{
			`The group chat discussion revolved around planning two outings
            - one for the IDP group and another to celebrate with a friend
            named Antoine`,
			`A supper outing at Swee Choon was also mentioned.`,
			`The conversation included logistical questions about fetching,
            sleeping schedules, and venue preferences.`,
			`Tasks were delegated for researching places and making reservations.`,
			`Two specific events were finalised with dates and times.`,
		},
		ImportantPoints: []string{
			`Prioritise CS202 meeting and CS420 research proposal.`,
			`Review CS206 code before Tuesday`,
			`Collaborate on CS205 exam prep.`,
			`Attend Samba Masala practices for upcoming open house gig.`,
		},
	},
}

func (server *Server) htmxHomePanel(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/htmx/home.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, nil)
}

func (server *Server) htmxTodoCard(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/htmx/todoCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"todos": fakeHomeToDoData,
	})
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
func (server *Server) htmxSummaryCard(writer http.ResponseWriter,
	request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/htmx/summaryCard.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{
		"summaries": fakeHomeSummaryData,
	})
}

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
func (server *Server) htmxSettings(writer http.ResponseWriter,
	request *http.Request) {
	tmpl, err := template.ParseFiles("./templates/htmx/settings.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(writer, map[string]interface{}{})

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
