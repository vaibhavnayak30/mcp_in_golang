package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vaibhavnayak30/mcp_in_golang/internal/controller/response"
)

// RoutesForConfig handles routing for configurable related resources 
func RoutesForConfig(r chi.Router) {
	r.Get("/", configGet)
}

// configGet returns the configurable information for the service 
// @Summary configGet
// @Description configGet
// @Accept json
// @Produce json
// @Success 200 {object} response.Config
// @Router /api/v1/config [get]
func configGet(w http.ResponseWriter, _ *http.Request) {
	conf := response.Config{
		AppInfo: &response.AppInfo{
			AppName: "Math MCP Server",
			Version: "1.0.0",
			GoVersion: "go1.24",
			BuildTime: "2024-0-01T00",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conf)
}