package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"gopkg.in/boj/redistore.v1"
)

type Server struct {
	listenAddr        string
	redisSessionStore redistore.RediStore
	router            *mux.Router
	telebotAddr       string
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
		router:            router,
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

	router.HandleFunc("/auth/{provider}", server.authHandler)
	router.HandleFunc("/auth/{provider}/callback", server.authCallback)
	router.HandleFunc("/logout/{provider}", server.logout)

	router.HandleFunc("/home", authWrapper(server.indexHome))
	router.HandleFunc("/bots", authWrapper(server.indexBots))
	router.HandleFunc("/settings", authWrapper(server.indexSettings))

	router.HandleFunc("/htmx/home", authWrapper(server.htmxHomePanel))
	router.HandleFunc("/htmx/home/todoCard", authWrapper(server.htmxTodoCard))
	router.HandleFunc("/htmx/home/eventCard", authWrapper(server.htmxEventCard))
	router.HandleFunc("/htmx/home/summaryCard", authWrapper(server.htmxSummaryCard))
	router.HandleFunc("/htmx/bots", authWrapper(server.htmxBots))
	router.HandleFunc("/htmx/botModal", authWrapper(server.htmxBotModal))
	router.HandleFunc("/htmx/settings", authWrapper(server.htmxSettings))

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
