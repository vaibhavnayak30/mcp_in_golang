// Package router provides routing configuration for maths MCP server
package router 

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}

// SetupRouter set routing for our services 
func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	// Setup api router 
	router.Route("api/v1", func(subRouter chi.Router) {
		// API 
	})
}