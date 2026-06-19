package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type ServerParams struct {
	// prefix name for logs
	Name string
	// addr of the http.server
	Addr string
	// http handler
	Handler http.Handler
	// catch shutdown signal from the os
	ShutdownSignalCtx context.Context
	// max time allowed for finishing everything, otherwise it will be cut.
	ShutdownGracePeriod time.Duration
}

// StartServer function runs the http.server with the specified params
func StartServer(params ServerParams) error {
	srv := &http.Server{
		Addr:    params.Addr,
		Handler: params.Handler,
	}

	errCh := make(chan error, 1)

	go func() {
		slog.Info("server started", "addr", params.Addr)

		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("listen and serve: %w", err)
			return
		}

		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err

	case <-params.ShutdownSignalCtx.Done():
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			params.ShutdownGracePeriod,
		)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown: %w", err)
		}

		return <-errCh
	}
}
