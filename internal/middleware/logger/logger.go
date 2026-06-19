package logger

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
)

// Logger functions configures router logs and control what to show and what to ignore
func Logger(log *slog.Logger, logFormat *httplog.Schema) func(next http.Handler) http.Handler {
	return httplog.RequestLogger(log, &httplog.Options{
		Level:  slog.LevelInfo,
		Schema: logFormat,

		RecoverPanics: true,
		Skip: func(req *http.Request, respStatus int) bool {
			return respStatus == 404 || respStatus == 405
		},
		LogRequestHeaders:  []string{"Origin"},
		LogResponseHeaders: []string{},
		LogExtraAttrs: func(req *http.Request, reqBody string, respStatus int) []slog.Attr {
			requestId := middleware.GetReqID(req.Context())

			return []slog.Attr{{Key: "id", Value: slog.StringValue(requestId)}}

		},
	})
}
