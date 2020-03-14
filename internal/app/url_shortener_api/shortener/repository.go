package shortener

import (
	"context"
	"url_shortener_api/internal/app/url_shortener_api/models"
)

type Repository interface {
	InsertUrl(ctx context.Context, url string, ttlSeconds int) (*models.Link, error)
	GetByUrl(ctx context.Context, url string) (*models.Link, error)
	UpdateLink(ctx context.Context, link *models.Link) (*models.Link, error)
	GetById(ctx context.Context, id int64, ttlSeconds int) (*models.Link, error)
	GetAllLinks(ctx context.Context) (*models.AllLinks, error)
	DeleteExpired() error
}
