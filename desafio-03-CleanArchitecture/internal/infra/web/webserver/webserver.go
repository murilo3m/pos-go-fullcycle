package webserver

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router chi.Router
	//Handlers      map[string]http.HandlerFunc
	Routes        []Route
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router: chi.NewRouter(),
		//Handlers:      make(map[string]http.HandlerFunc),
		Routes:        make([]Route, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Routes = append(s.Routes, Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	for _, route := range s.Routes {
		s.Router.Method(route.Method, route.Path, route.Handler)
	}

	if err := http.ListenAndServe(s.WebServerPort, s.Router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	//http.ListenAndServe(s.WebServerPort, s.Router)
}
