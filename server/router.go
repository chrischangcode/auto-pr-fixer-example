package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter creates and configures the application router.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handleRoot)
	r.Get("/health", handleHealth)

	return r
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
