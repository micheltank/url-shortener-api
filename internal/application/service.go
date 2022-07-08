package application

import (
	"context"
	"github.com/micheltank/url-shortener-api/internal/domain"
	"github.com/micheltank/url-shortener-api/internal/infra/repository"
	"github.com/pkg/errors"
)

type Service struct {
	repository domain.Repository
}

func NewService(repository *repository.UrlCassandraRepository) *Service {
	s := Service{
		repository: repository,
	}
	return &s
}

func (s *Service) Create(ctx context.Context, longUrl string) (*domain.Url, error) {
	url := domain.NewUrl(longUrl)
	err := s.repository.Create(ctx, longUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create alert on database")
	}
	return &url, nil
}

func (s *Service) Get(ctx context.Context, shorUrl string) (*domain.Url, error) {
	url, err := s.repository.Get(ctx, shorUrl)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (s *Service) Delete(ctx context.Context, shorUrl string) error {
	err := s.repository.Delete(ctx, shorUrl)
	if err != nil {
		return err
	}
	return nil
}
