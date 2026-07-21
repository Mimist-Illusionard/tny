package app

import (
	"github.com/Mimist-Illusionard/url-shortener/internal/config"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/Mimist-Illusionard/url-shortener/internal/transport/http"
)

func Run(cfg *config.Config) error {
	repo, err := repository.New(cfg)
	if err != nil {
		return err
	}

	return http.Serve(cfg, repo)
}
