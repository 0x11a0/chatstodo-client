package main

import (
    "html/template"
    "net/http"
)

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
