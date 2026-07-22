package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

type Handler struct {
	repo repository.Repository
}

func RegisterHandlers(repo repository.Repository) http.Handler {
	h := &Handler{repo: repo}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/links", h.CreateShortLink)
	mux.HandleFunc("GET /{shortCode}", h.GetShortLink)

	return mux
}

func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	req := CreateShortLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Errorf("incorrect request").Error(), http.StatusBadRequest)
		return
	}

	u := domain.NewUrl(req.URL)
	if err := h.repo.Create(*u); err != nil {
		http.Error(w, fmt.Errorf("error creating url").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&u)
}

func (h *Handler) GetShortLink(w http.ResponseWriter, r *http.Request) {
	short := r.PathValue("shortCode")
	if short == "" {
		http.Error(w, fmt.Errorf("missing short code").Error(), http.StatusBadRequest)
		return
	}

	u, err := h.repo.Get(short)
	if err != nil {
		http.Error(w, fmt.Errorf("error getting url %w").Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, u.Original, http.StatusSeeOther)
}
