package shortener

import (
	"context"
	"url_shortener_api/internal/app/url_shortener_api/models"
)

type Usecase interface {
	NewShortLink(ctx context.Context, url string) (*models.LinkResponse, error)
	Get(ctx context.Context, shortUrl string) (string, error)
	GetAll(ctx context.Context) (*models.AllLinks, error)
	SetTtl(ttl int) error
	Encode(id int64) string
	Decode(shortLink string) int64
}
