package grpc

import (
	"context"
	"errors"

	urlgrpc "github.com/Mimist-Illusionard/tny/api/v1/grpc"
	"github.com/Mimist-Illusionard/tny/internal/domain"
	"github.com/Mimist-Illusionard/tny/internal/repository"
	"github.com/Mimist-Illusionard/tny/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Shortener interface {
	CreateShortLink(ctx context.Context, originalURL string) (*domain.URL, error)
	GetOriginalLink(ctx context.Context, short string) (string, error)
}

type Handler struct {
	urlgrpc.UnimplementedURLShortenerServer
	service Shortener
}

func NewHandler(s Shortener) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateShortLink(ctx context.Context, req *urlgrpc.CreateShortLinkRequest) (*urlgrpc.CreateShortLinkResponse, error) {
	u, err := h.service.CreateShortLink(ctx, req.GetOriginalUrl())
	if err != nil {
		return nil, mapError(err)
	}

	return &urlgrpc.CreateShortLinkResponse{
		Id:          u.ID,
		ShortCode:   u.Short,
		OriginalUrl: u.Original,
	}, nil
}

func (h *Handler) GetOriginalLink(ctx context.Context, req *urlgrpc.GetOriginalLinkRequest) (*urlgrpc.GetOriginalLinkResponse, error) {
	originalURL, err := h.service.GetOriginalLink(ctx, req.GetShortCode())
	if err != nil {
		return nil, mapError(err)
	}

	return &urlgrpc.GetOriginalLinkResponse{OriginalUrl: originalURL}, nil
}

func mapError(err error) error {
	switch {
	case errors.Is(err, service.ErrInvalidURL), errors.Is(err, service.ErrInvalidShortCode):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, repository.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, service.ErrCannotGenerate):
		return status.Error(codes.ResourceExhausted, err.Error())
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
