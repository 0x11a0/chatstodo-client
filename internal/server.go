package internal

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

func InitServer(redisSessionStore *redistore.RediStore) *Server {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		log.Println("Optional env variable \"LISTEN_ADDR\" is not set. Defaulting to \"localhost:3000\".")
		listenAddr = "localhost:3000"
	}
	router := mux.NewRouter()
	server := &Server{
		listenAddr:        listenAddr,
		redisSessionStore: *redisSessionStore,
		router:            router,
		googleOAuthConfig: initGoogleOAuth(),
		oAuthVerifier:     oauth2.GenerateVerifier(),
	}

	log.Println("Server running on: http://" + listenAddr)
	return server
}

func (server *Server) Run() {
	CSRFSecret := os.Getenv("CSRF_SECRET")
	isProd := os.Getenv("IS_PROD")

	if CSRFSecret == "" {
		log.Fatalln("Required env variable \"CSRF_SECRET\" is not set. Exiting.")
	} else if isProd == "" {
		log.Fatalln("Required env variable \"IS_PROD\" is not set. Exiting.")
	}

	CSRF := csrf.Protect([]byte(CSRFSecret),
		csrf.Secure(isProd == "true"))

	router := server.router
	router.HandleFunc("/", server.authWrapper(server.dashboardHome))
	router.HandleFunc("/login", server.loginPage)

	router.HandleFunc("/auth/google", server.authHandler)
	router.HandleFunc("/auth/google/callback", server.googleAuthCallback)
	router.HandleFunc("/logout/google", server.logout)

	router.HandleFunc("/home", server.authWrapper(server.dashboardHome))
	//router.HandleFunc("/bots", server.authWrapper(server.dashboardBots))
	router.HandleFunc("/groups", server.authWrapper(server.dashboardGroups))
	router.HandleFunc("/settings", server.authWrapper(server.dashboardSettings))

	router.HandleFunc("/htmx/home", server.authWrapper(server.htmxHomePanel))
	router.HandleFunc("/htmx/home/tasks", server.authWrapper(server.htmxTasks))
	router.HandleFunc("/htmx/home/events", server.authWrapper(server.htmxEvents))
	router.HandleFunc("/htmx/home/summaries", server.authWrapper(server.htmxSummaries))
	//router.HandleFunc("/htmx/bots", server.authWrapper(server.htmxBots))
	//router.HandleFunc("/htmx/bots/modal", server.authWrapper(server.htmxBotModal))
	router.HandleFunc("/htmx/groups", server.authWrapper(server.htmxGroups))

	router.HandleFunc("/htmx/settings", server.authWrapper(server.htmxSettings))
	router.HandleFunc("/htmx/settings/platform", server.authWrapper(server.htmxSettingsPlatform))

	router.HandleFunc("/error/{code}", server.errorPageGeneric)
	router.NotFoundHandler = http.HandlerFunc(server.errorPageNotFound)

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
