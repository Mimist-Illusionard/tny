package app

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/Mimist-Illusionard/tny/internal/config"
	"github.com/Mimist-Illusionard/tny/internal/repository"
	"github.com/Mimist-Illusionard/tny/internal/repository/memory"
	"github.com/Mimist-Illusionard/tny/internal/repository/postgres"
	"github.com/Mimist-Illusionard/tny/internal/service"
	grpctransport "github.com/Mimist-Illusionard/tny/internal/transport/grpc"
	httptransport "github.com/Mimist-Illusionard/tny/internal/transport/http"
)

type closeFunc func() error

func Run(cfg *config.Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	return RunContext(ctx, cfg)
}

func RunContext(ctx context.Context, cfg *config.Config) error {
	repo, closeRepository, err := newRepository(cfg)
	if err != nil {
		return err
	}
	defer func() {
		_ = closeRepository()
	}()

	shortener := service.NewService(repo)
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, 2)
	go func() {
		errCh <- httptransport.Serve(runCtx, cfg.HTTPPort, shortener)
	}()
	go func() {
		errCh <- grpctransport.Serve(runCtx, cfg.GRPCPort, shortener)
	}()

	firstErr := <-errCh
	cancel()
	secondErr := <-errCh

	if firstErr != nil {
		return firstErr
	}
	if secondErr != nil {
		return secondErr
	}

	return nil
}

func newRepository(cfg *config.Config) (repository.Repository, closeFunc, error) {
	switch cfg.DBType {
	case config.Memory:
		return memory.New(), func() error { return nil }, nil
	case config.Postgres:
		db, err := postgres.Connect(cfg)
		if err != nil {
			return nil, nil, err
		}

		repo := postgres.New(db)
		return repo, repo.Close, nil
	default:
		return nil, nil, fmt.Errorf("invalid database type %q", cfg.DBType)
	}
}
