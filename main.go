package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aantoschuk/go-template/internal/middleware/logger"
	"github.com/aantoschuk/go-template/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	_, _ = logger.CreateLogger("dev", "api")

	serverParams := server.ServerParams{
		Addr:                "127.0.0.1:8000",
		Name:                "api",
		Handler:             nil,
		ShutdownSignalCtx:   ctx,
		ShutdownGracePeriod: 5 * time.Second,
	}

	err := server.StartServer(serverParams)
	if err != nil {
		panic(err)
	}

}
