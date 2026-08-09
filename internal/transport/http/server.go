package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const shutdownTimeout = 5 * time.Second

func Serve(ctx context.Context, port string, service Shortener) error {
	server := &http.Server{
		Addr:              ":" + port,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           RegisterHandlers(service),
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("http server listening on %s", server.Addr)
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		errCh <- err
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown http server: %w", err)
		}

		return <-errCh
	}
}
