package router

import (
	"github.com/go-chi/cors"
	"net/http"
)

var corsOptions = cors.Options{
	// allow request from every origin
	AllowedOrigins: []string{
		"https://*",
		"http://*",
	},
	AllowedMethods: []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	},
	AllowedHeaders: []string{
		"Accept",
		"Authorization",
		"Content-Type",
		"X-CSRF-Token",
	},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300,
}
