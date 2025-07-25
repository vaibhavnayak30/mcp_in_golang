// Package router provides routing configuration for maths MCP server
package router 

import (
	"net/http"

	"github.com/vaibhavnayak30/mcp_in_golang/internal/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// healthCheck godoc
// @Summary healthCheck
// @Description healthCehck
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/healthCheck [get]
func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("ok"))
}

// SetupRouter set routing for our services 
func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	// Add basic middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	// Setup api router 
	router.Route("api/v1", func(subRouter chi.Router) {
		// API documentation endpoint 
		subRouter.Get("/doc/*", httpSwagger.WrapHandler)

		// Health Check Endpoint 
		subRouter.Get("/healthCheck", healthCheck)

		// Configuration Endpoint 
		subRouter.Route("/config", controller.RoutesForConfig)

	})

	return router
}