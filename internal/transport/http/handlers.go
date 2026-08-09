package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/Mimist-Illusionard/url-shortener/internal/service"
)

const maxRequestBodySize = 1 << 20

type Shortener interface {
	CreateShortLink(ctx context.Context, originalURL string) (*domain.URL, error)
	GetOriginalLink(ctx context.Context, short string) (string, error)
}

type Handler struct {
	service Shortener
}

func RegisterHandlers(s Shortener) http.Handler {
	h := &Handler{service: s}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/links", h.CreateShortLink)
	mux.HandleFunc("GET /api/v1/links/{shortCode}", h.GetOriginalLink)
	mux.HandleFunc("GET /{shortCode}", h.RedirectToOriginalLink)

	return mux
}

func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	req := CreateShortLinkRequest{}
	if err := decoder.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	u, err := h.service.CreateShortLink(r.Context(), req.URL)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, CreateShortLinkResponse{
		ID:          u.ID,
		ShortCode:   u.Short,
		OriginalURL: u.Original,
	})
}

func (h *Handler) GetOriginalLink(w http.ResponseWriter, r *http.Request) {
	originalURL, err := h.service.GetOriginalLink(r.Context(), r.PathValue("shortCode"))
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, GetOriginalLinkResponse{OriginalURL: originalURL})
}

func (h *Handler) RedirectToOriginalLink(w http.ResponseWriter, r *http.Request) {
	originalURL, err := h.service.GetOriginalLink(r.Context(), r.PathValue("shortCode"))
	if err != nil {
		writeServiceError(w, err)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidURL), errors.Is(err, service.ErrInvalidShortCode):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, repository.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, service.ErrCannotGenerate):
		writeError(w, http.StatusServiceUnavailable, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
