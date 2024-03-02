package main

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
	"os"
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

func initAuth(cookieStore *redistore.RediStore) {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"),
			"http://localhost:3000/auth/google/callback"),
	)
	gothic.Store = cookieStore
}

func (server *Server) authCallback(writer http.ResponseWriter,
	request *http.Request) {
	user, err := gothic.CompleteUserAuth(writer, request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	log.Println(user)
	gothic.StoreInSession("userId", "testUser", request, writer)
	server.indexHome(writer, request)
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
	_, err := gothic.GetFromSession("userId", request)
	return err == nil
}
