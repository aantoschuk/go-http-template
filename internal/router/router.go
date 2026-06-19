package router

import (
	"log/slog"
	"net/http"

	loggerMiddleware "github.com/aantoschuk/go-template/internal/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
)

// struct to help handlers access needed services / values.
type CreateNewRouterParams struct {
	// route logger
	Logger *slog.Logger
	// log schema
	LogFormat *httplog.Schema
}

// CreateNewRouter return a new global router that should be attached to the
// transport layer.
func CreateNewRouter(params CreateNewRouterParams) http.Handler {
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.RequestID)
	r.Use(loggerMiddleware.Logger(params.Logger, params.LogFormat))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	return r

}
