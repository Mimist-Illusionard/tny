package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mimist-Illusionard/url-shortener/internal/service"
)

type Handler struct {
	s *service.Service
}

func RegisterHandlers(s *service.Service) http.Handler {
	h := &Handler{s: s}

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

	u, err := h.s.CreateShortLink(r.Context(), req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	link, err := h.s.GetOriginalLink(r.Context(), short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}
