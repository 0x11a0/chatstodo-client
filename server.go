package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"gopkg.in/boj/redistore.v1"
)

type Server struct {
	listenAddr        string
	redisSessionStore redistore.RediStore
	router            *mux.Router
	googleOAuthConfig *oauth2.Config
	oAuthVerifier     string
}

func initServer(redisSessionStore *redistore.RediStore) *Server {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		log.Panicln("Env variable \"LISTEN_ADDR\" has not been set. Exiting.")
	}
	router := mux.NewRouter()
	server := &Server{
		listenAddr:        listenAddr,
		redisSessionStore: *redisSessionStore,
		//gcAPI:             *gcAPI,
		router:            router,
		googleOAuthConfig: initAuth(),
		oAuthVerifier:     oauth2.GenerateVerifier(),
	}

	log.Println("Server running on: http://" + listenAddr)
	return server
}

func (server *Server) run() {
	CSRF := csrf.Protect([]byte(os.Getenv("CSRF_SECRET")),
		csrf.Secure(os.Getenv("IS_PROD") == "true"))

	router := server.router
	router.HandleFunc("/", server.indexHome)
	router.HandleFunc("/login", server.loginPage)

	router.HandleFunc("/auth/google", server.authHandler)
	router.HandleFunc("/auth/google/callback", server.authCallback)
	router.HandleFunc("/logout/google", server.logout)

	router.HandleFunc("/home", server.authWrapper(server.indexHome))
	router.HandleFunc("/bots", server.authWrapper(server.indexBots))
	router.HandleFunc("/settings", server.authWrapper(server.indexSettings))

	router.HandleFunc("/htmx/home", server.authWrapper(server.htmxHomePanel))
	router.HandleFunc("/htmx/home/todoCard", server.authWrapper(server.htmxTodoCard))
	router.HandleFunc("/htmx/home/eventCard", server.authWrapper(server.htmxEventCard))
	router.HandleFunc("/htmx/home/summaryCard", server.authWrapper(server.htmxSummaryCard))
	router.HandleFunc("/htmx/bots", server.authWrapper(server.htmxBots))
	router.HandleFunc("/htmx/botModal", server.authWrapper(server.htmxBotModal))
	router.HandleFunc("/htmx/settings", server.authWrapper(server.htmxSettings))

	router.HandleFunc("/api/addEvent", server.authWrapper(func(writer http.ResponseWriter,
		request *http.Request) {
		server.addEvent(writer, request)
	}))

	server.addFileServer()

	http.ListenAndServe(server.listenAddr, CSRF(server.router))
}

func writeJson(writer http.ResponseWriter, statusCode int, value any) error {
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(writer).Encode(value)
}

func (server *Server) addFileServer() {
	fileServer := http.FileServer(http.Dir("./static"))
	server.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
}
