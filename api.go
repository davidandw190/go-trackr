package gotrackr

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// APIServer represents the HTTP API server.
type APIServer struct {
	addr  string
	store Storage
}

// NewAPIServer creates a new APIServer instance.
func NewAPIServer(addr string, store Storage) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// Serve starts the API server.
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Register services
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting API server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
