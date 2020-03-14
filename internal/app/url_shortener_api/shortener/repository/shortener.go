package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"url_shortener_api/internal/app/url_shortener_api/models"
	"url_shortener_api/internal/app/url_shortener_api/shortener"
)

type shortenerRepository struct {
	db *sql.DB
}

func (s shortenerRepository) DeleteExpired() error {
	fmt.Println("Cron running", time.Now())

	_, err := s.db.Exec("DELETE from links WHERE expires < now()")

	if err != nil {
		fmt.Println("Cron failed")
	}

	return err
}

func (s shortenerRepository) GetAllLinks(ctx context.Context) (*models.AllLinks, error) {
	links := &models.AllLinks{}
	links.Result = make(models.LinkPairs, 0)
	//links.Result = make([]models.Link, 0)

	rows, err := s.db.QueryContext(
		ctx,
		"SELECT url, short_url FROM links",
		//"SELECT id, url, short_url, created, expires FROM links",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		l := models.LinkPair{}
		//l := models.Link{}
		err := rows.Scan(&l[0], &l[1])
		//err := rows.Scan(&l.Id, &l.Url, &l.Short_url, &l.Created, &l.Expires)

		if err != nil {
			return nil, err
		}

		links.Result = append(links.Result, l)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	links.Total = int64(len(links.Result))

	return links, err
}

func (s shortenerRepository) GetById(ctx context.Context, id int64, ttlSeconds int) (*models.Link, error) {
	link := &models.Link{}
	ttl := time.Now().Local().Add(time.Duration(ttlSeconds) * time.Second)

	err := s.db.QueryRowContext(
		ctx,
		"UPDATE links SET expires = $2 "+
			"WHERE id = $1 AND expires > now() RETURNING id, url, short_url, created, expires",
		id,
		ttl,
	).Scan(
		&link.Id,
		&link.Url,
		&link.Short_url,
		&link.Created,
		&link.Expires,
	)

	if err != nil {
		return nil, err
	}

	return link, err
}

func (s shortenerRepository) InsertUrl(ctx context.Context, url string, ttlSeconds int) (*models.Link, error) {
	link := &models.Link{}
	ttl := time.Now().Local().Add(time.Duration(ttlSeconds) * time.Second)

	err := s.db.QueryRowContext(
		ctx,
		"INSERT INTO links (url, expires) "+
			"VALUES ($1, $2) RETURNING id, url, created",
		url,
		ttl,
	).Scan(
		&link.Id,
		&link.Url,
		&link.Created,
	)

	if err != nil {
		return nil, err
	}

	return link, err
}

func (s shortenerRepository) GetByUrl(ctx context.Context, url string) (*models.Link, error) {
	link := &models.Link{}

	err := s.db.QueryRowContext(
		ctx,
		"SELECT id, url, short_url, created, expires "+
			"FROM links WHERE url = $1",
		url,
	).Scan(
		&link.Id,
		&link.Url,
		&link.Short_url,
		&link.Created,
		&link.Expires,
	)

	if err != nil {
		return nil, err
	}

	return link, err
}

func (s shortenerRepository) UpdateLink(ctx context.Context, link *models.Link) (*models.Link, error) {
	err := s.db.QueryRowContext(
		ctx,
		"UPDATE links SET short_url = $2"+
			" WHERE id = $1 RETURNING id, url, short_url, created, expires",
		link.Id,
		link.Short_url,
	).Scan(
		&link.Id,
		&link.Url,
		&link.Short_url,
		&link.Created,
		&link.Expires,
	)

	if err != nil {
		return nil, err
	}

	return link, err
}

func NewShortenerRepository(db *sql.DB) shortener.Repository {
	return &shortenerRepository{db}
}
