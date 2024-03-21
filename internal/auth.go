package internal

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"time"
)

// Constants for cookie naming scheme
const (
	COOKIE_NAME                 = "ctd-cookie"
	COOKIE_GOOGLE_TOKEN_SOURCE  = "google_token_source"
	COOKIE_GOOGLE_ACCESS_TOKEN  = "google_access_token"
	COOKIE_GOOGLE_REFRESH_TOKEN = "google_refresh_token"
	COOKIE_GOOGLE_EXPIRES_AT    = "google_expires_at"
	COOKIE_EMAIL                = "email"
	COOKIE_JWT                  = "jwt"
)

// Function prototype for the authWrapper below
type serverFunc func(http.ResponseWriter, *http.Request)

// Wraps any http.HandleFunc functions. Requires the
// browser to be logged in, else defaults to login page.
// Used for ALL possible routes that are exposed.
func (server *Server) authWrapper(function serverFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !server.isValidSession(request) {
			http.Redirect(writer, request, "/login", http.StatusSeeOther)
		} else {
			function(writer, request)
		}
	}
}

var (
	GOOGLE_SCOPES = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/calendar.events",
	}
)

// Returns the oauth2.Config for Google. Requires
// GOOGLE_KEY, GOOGLE_SECRET, GOOGLE_CALLBACK_URL
// env variables to be set prior to calling this function.
func initGoogleOAuth() *oauth2.Config {
	clientID := os.Getenv("GOOGLE_KEY")
	clientSecret := os.Getenv("GOOGLE_SECRET")
	redirectURL := os.Getenv("GOOGLE_CALLBACK_URL")
	if clientID == "" {
		log.Fatalln("Required env variable \"GOOGLE_KEY\" is not set. Exiting.")
	} else if clientSecret == "" {
		log.Fatalln("Required env variable \"GOOGLE_SECRET\" is not set. Exiting.")
	} else if redirectURL == "" {
		log.Fatalln("Required env variable \"GOOGLE_CALLBACK_URL\" is not set. Exiting.")
	}

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     google.Endpoint,
		Scopes:       GOOGLE_SCOPES,
	}
}

// /auth/google/callback
// Auth callback function for Google oauth.
// Any errors during callback will redirect back to login
func (server *Server) googleAuthCallback(writer http.ResponseWriter,
	request *http.Request) {
	googleOAuthCode := request.FormValue("code")
	token, err := server.googleOAuthConfig.Exchange(
		request.Context(),
		googleOAuthCode,
		oauth2.VerifierOption(server.oAuthVerifier),
	)
	if err != nil {
		log.Println("auth.go - authCallback(), get token")
		log.Println(err)
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	session, err := server.redisSessionStore.Get(
		request,
		COOKIE_NAME,
	)
	if err != nil {
		log.Println("auth.go - authCallback(), get session")
		log.Println(err)
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(
		token.Extra("id_token").(string),
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("GOOGLE_SECRET")), nil
		},
	)

	session.Values[COOKIE_GOOGLE_ACCESS_TOKEN] = token.AccessToken
	session.Values[COOKIE_GOOGLE_REFRESH_TOKEN] = token.RefreshToken
	session.Values[COOKIE_GOOGLE_EXPIRES_AT] = token.Expiry.Format(time.RFC3339)
	session.Values[COOKIE_EMAIL] = claims["email"]

	err = session.Save(request, writer)
	if err != nil {
		log.Println("auth.go - authCallback(), save session")
		log.Println(err)
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	log.Println("authCallback success")
	server.dashboardHome(writer, request)
}

// Returns the token if found, or nil if error
func (server *Server) getGoogleOAuthToken(request *http.Request) *oauth2.Token {
	session, err := server.redisSessionStore.Get(
		request,
		COOKIE_NAME,
	)
	if err != nil {
		log.Println("auth.go - getGooleOAuthToken() get session")
		log.Println(err)
		return nil
	}

	googleExpiresAt := session.Values[COOKIE_GOOGLE_EXPIRES_AT]

	expireTime, err := time.Parse(time.RFC3339, googleExpiresAt.(string))
	if err != nil {
		log.Println("auth.go - getGooleOAuthToken(), parse time")
		log.Println(err)
		return nil
	}

	return &oauth2.Token{
		AccessToken:  session.Values[COOKIE_GOOGLE_ACCESS_TOKEN].(string),
		RefreshToken: session.Values[COOKIE_GOOGLE_REFRESH_TOKEN].(string),
		Expiry:       expireTime,
		TokenType:    "Bearer",
	}
}

// Returns the token source if found, or nil if error
func (server *Server) getGoogleOAuthTokenSource(request *http.Request) oauth2.TokenSource {
	googleOAuthToken := server.getGoogleOAuthToken(request)
	if googleOAuthToken == nil {
		log.Println("no token found")
		return nil
	}
	return server.googleOAuthConfig.TokenSource(
		request.Context(),
		googleOAuthToken,
	)
}

// Refreshes the google access token if its within 5 minutes
// of expiring
func (server *Server) refreshGoogleAccessToken(writer http.ResponseWriter,
	request *http.Request) error {
	// Will never error here since this function is wrapped with the auth wrapper
	googleOAuthToken := server.getGoogleOAuthToken(request)
	if googleOAuthToken.Expiry.After(time.Now().Add(5 * time.Minute)) {
		log.Println("No need to refresh token")
		return nil

	}

	log.Println("token expiring soon, refreshing")
	log.Println("oldToken", googleOAuthToken.AccessToken)
	newToken, err := server.googleOAuthConfig.TokenSource(
		request.Context(),
		googleOAuthToken,
	).Token()
	log.Println("newToken", newToken.AccessToken)
	session, err := server.redisSessionStore.Get(
		request,
		COOKIE_NAME,
	)
	if err != nil {
		log.Println("auth.go - refreshGoogleAccessToken() get session")
		log.Println(err)
		return nil
	}
	session.Values[COOKIE_GOOGLE_ACCESS_TOKEN] = newToken.AccessToken
	session.Save(request, writer)
	return nil
}

// /logout/google/
func (server *Server) logout(writer http.ResponseWriter,
	request *http.Request) {
	session, err := server.redisSessionStore.Get(
		request,
		COOKIE_NAME,
	)
	if err != nil {
		log.Println("auth.go - logout() get session")
		log.Println(err)
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(request, writer)
	if err != nil {
		log.Println("auth.go - logout(), save session")
		log.Println(err)
		return
	}

	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}

// /auth/google
func (server *Server) authHandler(writer http.ResponseWriter,
	request *http.Request) {
	if server.isValidSession(request) {
		server.dashboardHome(writer, request)
		return
	}

	url := server.googleOAuthConfig.AuthCodeURL(
		"state",
		oauth2.AccessTypeOffline,
		oauth2.S256ChallengeOption(server.oAuthVerifier),
	)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func (server *Server) isValidSession(request *http.Request) bool {
	session, err := server.redisSessionStore.Get(request, COOKIE_NAME)
	if err != nil {
		log.Println("auth.go - isValidSession, getSession")
		log.Println(err)
		return false
	}
	log.Println(COOKIE_EMAIL, session.Values[COOKIE_EMAIL])
	return session.Values[COOKIE_EMAIL] != nil
}
