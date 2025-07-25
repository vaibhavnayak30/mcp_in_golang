// Package response defines response structure for HTTP API endpoints 
package response

// AppInfo represents application information 
type AppInfo struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	GoVersion string `json:"go_version"`
	BuildTime string `json:"build_time"`
}

// Config represents the configuration response structure 
type Config struct {
	AppInfo *AppInfo `json:"app_info"`
}