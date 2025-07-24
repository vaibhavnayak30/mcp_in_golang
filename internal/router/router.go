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