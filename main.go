package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aantoschuk/go-template/internal/middleware/logger"
	"github.com/aantoschuk/go-template/internal/router"
	"github.com/aantoschuk/go-template/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	logger, logFormat := logger.CreateLogger("dev", "api")

	r := router.CreateNewRouter(router.CreateNewRouterParams{
		Logger:    logger,
		LogFormat: logFormat,
	})

	serverParams := server.ServerParams{
		Addr:                "127.0.0.1:8000",
		Name:                "api",
		Handler:             r,
		ShutdownSignalCtx:   ctx,
		ShutdownGracePeriod: 5 * time.Second,
	}

	err := server.StartServer(serverParams)
	if err != nil {
		panic(err)
	}

}
