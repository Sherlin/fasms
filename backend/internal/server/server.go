package server

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
    h "github.com/gorilla/handlers"
	"fasms/internal/handlers"
	"fasms/internal/config"
	"fasms/internal/db"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Server represents the HTTP server.
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new Server instance.
func NewServer(cfg *config.Config) *Server {

	err := db.NewDB()
    if err != nil {
        log.Fatal("Error connecting to database: %v", err)
    }
    

	router := mux.NewRouter()
	corsHandler := h.CORS(
		h.AllowedOrigins([]string{"http://localhost:3000"}), // Update allowed origins as needed
		h.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),                    // Allow required HTTP methods
		h.AllowedHeaders([]string{"Content-Type", "Authorization"}),                              // Allow necessary headers
		h.AllowCredentials(),                                                                    // Allow cookies/auth headers
	)
	// Register routes
	router.HandleFunc("/health", handlers.HealthCheckHandler)
    router.HandleFunc("/api/applicants", handlers.CreateApplicant).Methods("POST")
    router.HandleFunc("/api/applicants", handlers.GetApplicants).Methods("GET")
	router.HandleFunc("/api/applicants/{id}", handlers.GetApplicantByID).Methods("GET")
    router.HandleFunc("/api/applicants/{id}", handlers.UpdateApplicant).Methods("PUT")
    router.HandleFunc("/api/applicants/{id}", handlers.DeleteApplicant).Methods("DELETE")
    router.HandleFunc("/api/schemes", handlers.CreateScheme).Methods("POST")
    router.HandleFunc("/api/schemes", handlers.GetSchemes).Methods("GET")
	router.HandleFunc("/api/schemes/eligible", handlers.GetSchemesForApplicant).Methods("GET")
    router.HandleFunc("/api/schemes/{id}", handlers.UpdateScheme).Methods("PUT")
    router.HandleFunc("/api/schemes/{id}", handlers.DeleteScheme).Methods("DELETE")
    router.HandleFunc("/api/applications", handlers.CreateApplication).Methods("POST")
    router.HandleFunc("/api/applications", handlers.GetApplications).Methods("GET")
    router.HandleFunc("/api/applications/{id}", handlers.UpdateApplication).Methods("PUT")
    router.HandleFunc("/api/applications/{id}", handlers.DeleteApplication).Methods("DELETE")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	
	// Create HTTP server
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      corsHandler(router),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{httpServer: httpServer}
}

// Run starts the server.
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}