package http

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mimist-Illusionard/url-shortener/internal/config"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

func Serve(cfg *config.Config, r repository.Repository) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:         ":" + "8080",
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	log.Printf("http server listening on port :8080")
	return s.ListenAndServe()
}
