package logger

import (
	"log/slog"
	"os"

	"github.com/go-chi/httplog/v3"
	"github.com/go-chi/traceid"
	"github.com/golang-cz/devslog"
)

// Create function rewrites slog logger
func CreateLogger(env, service string) (*slog.Logger, *httplog.Schema) {
	isLocal := env == "dev"

	logFormat := httplog.SchemaECS.Concise(isLocal)

	// depending on the environment it would display different info.
	// for example on dev it would have pretty json data, like service name and addr
	// while on prod it would show just long message with where and the function name
	// which is output from
	logger := slog.New(logHandler(isLocal, &slog.HandlerOptions{
		AddSource:   !isLocal,
		ReplaceAttr: logFormat.ReplaceAttr,
	})).With("service", service)

	// if it's not a dev add additional information such as a project name, version and env
	if !isLocal {
		logger = logger.With(
			slog.String("app", "go-template"),
			slog.String("version", "0.0.1"),
			slog.String("env", env),
		)
	}

	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	return logger, logFormat

}

func logHandler(isLocal bool, handlerOptions *slog.HandlerOptions) slog.Handler {
	if isLocal {
		return devslog.NewHandler(os.Stdout, &devslog.Options{
			SortKeys:           true,
			MaxErrorStackTrace: 5,
			MaxSlicePrintSize:  20,
			HandlerOptions:     handlerOptions,
		})
	}

	return traceid.LogHandler(
		slog.NewJSONHandler(os.Stdout, handlerOptions),
	)
}
