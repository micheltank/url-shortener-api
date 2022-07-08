package handler

import (
	"context"
	"github.com/micheltank/url-shortener-api/internal/domain"
	"github.com/micheltank/url-shortener-api/internal/port/grpc/pb"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service interface {
	Create(ctx context.Context, longUrl string) (*domain.Url, error)
	Get(ctx context.Context, shorUrl string) (*domain.Url, error)
	Delete(ctx context.Context, shorUrl string) error
}

type UrlHandler struct {
	pb.UnimplementedUrlHandlerServer
	service Service
}

func NewUrlHandler(service Service) *UrlHandler {
	return &UrlHandler{
		service: service,
	}
}

func (h *UrlHandler) Create(ctx context.Context, in *pb.LongUrlParam) (*pb.ShortUrl, error) {
	url, err := h.service.Create(ctx, in.GetLongUrl())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get urls at price")
	}
	return &pb.ShortUrl{
		Url: url.GetShortUrl(),
	}, nil
}

func (h *UrlHandler) Get(ctx context.Context, in *pb.ShortUrlParam) (*pb.LongUrl, error) {
	url, err := h.service.Get(ctx, in.GetShortUrl())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get urls at price")
	}
	return &pb.LongUrl{
		Url: url.GetLongUrl(),
	}, nil
}

func (h *UrlHandler) Delete(ctx context.Context, in *pb.ShortUrlParam) (*emptypb.Empty, error) {
	err := h.service.Delete(ctx, in.GetShortUrl())
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete url")
	}
	return &emptypb.Empty{}, nil
}
