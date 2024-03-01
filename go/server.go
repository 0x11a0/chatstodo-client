package main

import (
	"encoding/json"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	listenAddr        string
	redisSessionStore RedisSessionStore
	router            *mux.Router
	telebotAddr       string
}

func initServer(redisSessionStore *RedisSessionStore) *Server {

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
	router.HandleFunc("/home", server.indexHome)
	router.HandleFunc("/bots", server.indexBots)
	router.HandleFunc("/settings", server.indexSettings)
	router.HandleFunc("/htmx/homeCards", server.htmxHomeCards)
	router.HandleFunc("/htmx/bots", server.htmxBots)
	router.HandleFunc("/htmx/settings", server.htmxSettings)
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
