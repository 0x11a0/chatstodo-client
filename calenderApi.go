package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/markbates/goth/gothic"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarAPI struct {
	config *oauth2.Config
}

func initGoogleCalendarApi() *GoogleCalendarAPI {

	googleKey := os.Getenv("GOOGLE_KEY")
	googleSecret := os.Getenv("GOOGLE_SECRET")
	googleCallbackUrl := os.Getenv("GOOGLE_CALLBACK_URL")
	if googleKey == "" {
		log.Fatalln("GOOGLE_KEY not set, exiting.")
	} else if googleSecret == "" {
		log.Fatalln("GOOGLE_SECRET not set, exiting.")
	} else if googleCallbackUrl == "" {
		log.Fatalln("GOOGLE_CALLBACK_URL not set, exiting.")
	}

	conf := &oauth2.Config{
		ClientID:     googleKey,
		ClientSecret: googleSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  googleCallbackUrl,
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
	}

	return &GoogleCalendarAPI{
		config: conf,
	}
}

func (gcAPI *GoogleCalendarAPI) getCalendars(writer http.ResponseWriter,
	request *http.Request) {
	authURL := gcAPI.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Println("authURL", authURL)

	// NOTE: No need to check for err as it will be wrapped with auth wrapper
	refreshToken, _ := gothic.GetFromSession("googleRefreshToken", request)
	accessToken, _ := gothic.GetFromSession("googleAccessToken", request)
	accessExpiry, _ := gothic.GetFromSession("googleAccessExpiry", request)

	expiresAt, err := time.Parse(time.RFC3339, accessExpiry)
	if err != nil {
		log.Println("calendarApi.go - getClient(), expiresAt")
		log.Println(err)
		return
	}

	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       expiresAt,
	}

	googleClient := oauth2.NewClient(context.Background(),
		oauth2.StaticTokenSource(token))

	calendarService, err := calendar.NewService(context.Background(),
		option.WithHTTPClient(googleClient))
	if err != nil {
		log.Println("calendarApi.go - getClient(), calendarService")
		log.Println(err)
		return
	}

	calendarList, err := calendarService.CalendarList.List().Do()
	if err != nil {
		log.Println("calendarApi.go - getClient(), calendarList")
		log.Println(err)
		return
	}

	for _, item := range calendarList.Items {
		log.Println(item.Summary)
	}

}

func (gcAPI *GoogleCalendarAPI) addEvent(writer http.ResponseWriter,
	request *http.Request) {
	authURL := gcAPI.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Println("authURL", authURL)

	// NOTE: No need to check for err as it will be wrapped with auth wrapper
	refreshToken, _ := gothic.GetFromSession("googleRefreshToken", request)
	accessToken, _ := gothic.GetFromSession("googleAccessToken", request)
	accessExpiry, _ := gothic.GetFromSession("googleAccessExpiry", request)

	expiresAt, err := time.Parse(time.RFC3339, accessExpiry)
	if err != nil {
		log.Println("calendarApi.go - addEvent(), expiresAt")
		log.Println(err)
		return
	}

	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       expiresAt,
	}

	googleClient := oauth2.NewClient(context.Background(),
		oauth2.StaticTokenSource(token))

	event := &calendar.Event{
		Summary:     "test event",
		Location:    "Singapore!!",
		Description: "booya",
		Start: &calendar.EventDateTime{
			DateTime: "2024-03-28T09:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
	}

	calendarService, err := calendar.NewService(context.Background(),
		option.WithHTTPClient(googleClient))
	if err != nil {
		log.Println("calendarApi.go - addEvent(), calendarService")
		log.Println(err)
		return
	}
	calendarId := "primary"
	event, err = calendarService.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Println("calendarApi.go - addEvent(), do")
		log.Println(err)
		return

	}

	/*
		calendarList, err := calendarService.CalendarList.List().Do()
		if err != nil {
			log.Println("calendarApi.go - getClient(), calendarList")
			log.Println(err)
			return
		}

		for _, item := range calendarList.Items {
			log.Println(item.Summary)
		}
	*/
}
