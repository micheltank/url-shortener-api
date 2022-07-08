package domain

import "context"

type Repository interface {
	Get(ctx context.Context, shortUrl string) (Url, error)
	Create(ctx context.Context, longUrl string) error
	Delete(ctx context.Context, shortUrl string) error
}
