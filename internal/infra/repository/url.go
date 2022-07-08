package repository

import (
	"context"
	"github.com/micheltank/url-shortener-api/internal/domain"
)

type UrlCassandraRepository struct {
}

func NewUrlCassandraRepository() *UrlCassandraRepository {
	return &UrlCassandraRepository{}
}

func (u UrlCassandraRepository) Get(ctx context.Context, shortUrl string) (domain.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (u UrlCassandraRepository) Create(ctx context.Context, longUrl string) error {
	//TODO implement me
	panic("implement me")
}

func (u UrlCassandraRepository) Delete(ctx context.Context, shortUrl string) error {
	//TODO implement me
	panic("implement me")
}
