package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"gopkg.in/boj/redistore.v1"
)

type serverFunc func(http.ResponseWriter, *http.Request)

func authWrapper(function serverFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !isValidSession(request) {
			http.Redirect(writer, request, "/login", http.StatusSeeOther)
		} else {
			function(writer, request)
		}
	}
}

const GOOGLE_CALLBACK_URL = "http://localhost:3000/auth/google/callback"

func initAuth(cookieStore *redistore.RediStore) {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"),
			GOOGLE_CALLBACK_URL),
	)
	gothic.Store = cookieStore
}

func (server *Server) authCallback(writer http.ResponseWriter,
	request *http.Request) {
	user, err := gothic.CompleteUserAuth(writer, request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	if user.Provider == "" {
		server.indexHome(writer, request)
		return
	}
	log.Println("provider", user.Provider)
	log.Println("email", user.Email)
	if user.Provider == "google" {
		log.Println("googleAccessExpiry", user.ExpiresAt.Format(time.RFC3339))
		gothic.StoreInSession("googleRefreshToken", user.RefreshToken, request, writer)
		gothic.StoreInSession("googleAccessExpiry", user.ExpiresAt.Format(time.RFC3339),
			request, writer)
		gothic.StoreInSession("googleAccessToken", user.AccessToken, request, writer)
	}
	gothic.StoreInSession("email", user.Email, request, writer)

	server.indexHome(writer, request)
}

func (server *Server) refreshGoogleAccessToken(writer http.ResponseWriter,
	request *http.Request) error {
	// Will never error here since this function is wrapped with the auth wrapper
	accessExpiry, err := gothic.GetFromSession("googleAccessExpiry", request)

	accessTime, err := time.Parse(time.RFC3339, accessExpiry)
	if err != nil {
		log.Println("auth.go - refreshGoogleAccessToken()")
		return err

	} else if accessTime.After(time.Now().Add(5 * time.Minute)) {
		// Only refresh if within 5 minutes of expiry

		log.Println("No need to refresh token")
		return nil
	}

	log.Println("token expiring soon, refreshing")

	refreshToken, err := gothic.GetFromSession("googleRefreshToken", request)
	googleProvider, err := goth.GetProvider("google")
	if err != nil {
		log.Println("auth.go - refreshGoogleAccessToken()")
		return err
	}

	accessToken, err := googleProvider.RefreshToken(refreshToken)
	if err != nil {
		log.Println("auth.go - refreshGoogleAccessToken()")
		return err
	}
	log.Println("accessToken", accessToken)

	gothic.StoreInSession("googleAccessTime", fmt.Sprint(time.Now().UnixMilli()),
		request, writer)
	gothic.StoreInSession("googleAccessToken", accessToken.AccessToken, request, writer)
	return nil
}

func (server *Server) logout(writer http.ResponseWriter,
	request *http.Request) {
	gothic.Logout(writer, request)
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}

func (server *Server) authHandler(writer http.ResponseWriter,
	request *http.Request) {
	_, err := gothic.CompleteUserAuth(writer, request)
	if err != nil {
		gothic.BeginAuthHandler(writer, request)
		return
	}
	server.indexHome(writer, request)
}

func isValidSession(request *http.Request) bool {
	_, err := gothic.GetFromSession("email", request)
	return err == nil
}
