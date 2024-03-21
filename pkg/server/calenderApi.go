package server

import (
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func (server *Server) testAddEvent(writer http.ResponseWriter,
	request *http.Request) {
	event := calendar.Event{
		Summary:     "test event",
		Location:    "Singapore!!",
		Description: "booya",
		Start: &calendar.EventDateTime{
			DateTime: "2024-03-28T09:00:00-07:00",
			TimeZone: "Asia/Singapore",
		},
		End: &calendar.EventDateTime{
			DateTime: "2024-03-29T09:00:00-07:00",
			TimeZone: "Asia/Singapore",
		},
		Recurrence: []string{
			/*
				"RRULE:FREQ=DAILY;COUNT=2",
			*/
		},
		Attendees: []*calendar.EventAttendee{
			/*
				&calendar.EventAttendee{Email: "lpage@example.com"},
				&calendar.EventAttendee{Email: "sbrin@example.com"},
			*/
		},
	}
	server.addEvent(request, &event)
}


func (server *Server) addEvent(request *http.Request,
	event *calendar.Event) {

	googleOAuthToken := server.getGoogleOAuthToken(request)
	if googleOAuthToken == nil {
		log.Println("token missing")
		return
	}

	googleClient := server.googleOAuthConfig.Client(
		request.Context(),
		googleOAuthToken,
	)

	googleCalendarService, err := calendar.NewService(
		request.Context(),
		option.WithHTTPClient(googleClient))
	if err != nil {
		log.Println("calendarApi.go - addEvent(), calendarService")
		log.Println(err)
		return
	}

	calendarId := "primary"
	event, err = googleCalendarService.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Println("calendarApi.go - addEvent(), do")
		log.Println(err)
		return
	}
	log.Println("add success")
}
