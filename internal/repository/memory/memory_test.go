package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

func TestCreate(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	url := domain.NewURL("test", "t")
	err = r.Create(context.Background(), url)
	if err != nil {
		t.Fatalf("Create() unexpected error: %v", err)
	}

	err = r.Create(context.Background(), url)
	if !errors.Is(err, repository.ErrExists) {
		t.Fatalf("exist url: %v", err)
	}

	url = domain.NewURL("test", "unique")
	err = r.Create(context.Background(), url)
	if !errors.Is(err, repository.ErrNotUnique) {
		t.Fatalf("not unique url: %v", err)
	}

	if len(r.shortUrls) < 1 {
		t.Fatalf("short url len < 1: %v", r.shortUrls)
	}

	if len(r.origUrls) < 1 {
		t.Fatalf("short url len < 1: %v", r.shortUrls)
	}

	if _, ok := r.shortUrls["t"]; !ok {
		t.Fatalf("short url not created: %v", url)
	}

	if _, ok := r.origUrls[url.Original]; !ok {
		t.Fatalf("original url not created: %v", url)
	}
}

func TestGet(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	url := domain.NewURL("test", "t")
	err = r.Create(context.Background(), url)
	if err != nil {
		t.Fatalf("Create() unexpected error: %v", err)
	}

	_, err = r.Get(context.Background(), url.Short)
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}

	_, err = r.Get(context.Background(), "not exists")
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("not found err: %v", err)
	}
}

func TestGetByOriginal(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	url := domain.NewURL("test", "t")
	err = r.Create(context.Background(), url)
	if err != nil {
		t.Fatalf("Create() unexpected error: %v", err)
	}

	_, err = r.GetByOriginalURL(context.Background(), url.Original)
	if err != nil {
		t.Fatalf("GetByOriginalURL() unexpected error: %v", err)
	}

	_, err = r.GetByOriginalURL(context.Background(), "not exists")
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("not found err: %v", err)
	}
}

func TestDelete(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	url := domain.NewURL("test", "t")
	err = r.Create(context.Background(), url)
	if err != nil {
		t.Fatalf("Create() unexpected error: %v", err)
	}

	r.Delete(context.Background(), "t")

	if len(r.shortUrls) > 1 {
		t.Fatalf("short url len > 1: %v", r.shortUrls)
	}

	if len(r.origUrls) > 1 {
		t.Fatalf("orig url len > 1: %v", r.origUrls)
	}
}
