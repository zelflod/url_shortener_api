package usecase

import (
	"bytes"
	"context"
	"strings"
	"time"
	"url_shortener_api/internal/app/url_shortener_api/models"
	"url_shortener_api/internal/app/url_shortener_api/shortener"
)

type shortenerUsecase struct {
	shortenerRepo  shortener.Repository
	contextTimeout time.Duration
	alphabet       []string
	ttl            int
}

func (s *shortenerUsecase) SetTtl(ttl int) error {
	s.ttl = ttl

	return nil
}

func (s *shortenerUsecase) GetAll(ctx context.Context) (*models.AllLinks, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	links, err := s.shortenerRepo.GetAllLinks(ctx)

	if err != nil {
		return nil, err
	}

	return links, nil
}

func (s *shortenerUsecase) Get(ctx context.Context, shortUrl string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	id := s.Decode(shortUrl)
	link, err := s.shortenerRepo.GetById(ctx, id, s.ttl)

	if err != nil {
		return "", err
	}

	return link.Url, nil
}

func (s *shortenerUsecase) NewShortLink(ctx context.Context, url string) (*models.LinkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	link, err := s.shortenerRepo.InsertUrl(ctx, url, s.ttl)
	if err != nil {
		link, err = s.shortenerRepo.GetByUrl(ctx, url)
		if err != nil {
			return nil, err
		}
	}

	link.Short_url = s.Encode(link.Id)

	link, err = s.shortenerRepo.UpdateLink(ctx, link)
	if err != nil {
		return nil, err
	}

	shortLink := &models.LinkResponse{Link: link.Short_url}

	return shortLink, nil
}

func NewShortenerUsecase(s shortener.Repository, timeout time.Duration, ttl int) shortener.Usecase {
	return &shortenerUsecase{
		shortenerRepo:  s,
		contextTimeout: timeout,
		alphabet:       strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", ""),
		ttl:            ttl,
	}
}

func (s *shortenerUsecase) Encode(id int64) string {
	var b bytes.Buffer
	var radix int64 = 62
	for id > 0 {
		b.WriteString(s.alphabet[id%radix])

		id = id / radix
	}

	return reverse(b.String())
}

func (s *shortenerUsecase) Decode(shortLink string) int64 {
	var id int64 = 0
	var radix int64 = 62

	for i := 0; i < len(shortLink); i++ {
		if 'a' <= shortLink[i] && shortLink[i] <= 'z' {
			id = id*radix + int64(shortLink[i]) - 'a'
		}

		if 'A' <= shortLink[i] && shortLink[i] <= 'Z' {
			id = id*radix + int64(shortLink[i]) - 'A' + 26
		}

		if '0' <= shortLink[i] && shortLink[i] <= '9' {
			id = id*radix + int64(shortLink[i]) - '0' + 52
		}
	}

	return id
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
