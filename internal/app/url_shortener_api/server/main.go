package server

import (
	"database/sql"
	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo/v4"
	"log"
	"time"
	"url_shortener_api/internal/app/url_shortener_api/shortener/delivery/http"
	"url_shortener_api/internal/app/url_shortener_api/shortener/repository"
	"url_shortener_api/internal/app/url_shortener_api/shortener/usecase"
	db2 "url_shortener_api/internal/db"
)

type Config struct {
	DatabaseURL  string
	Address      string
	Timeout      int
	Ttl          int
	CronInterval int
}

func StartApp(config Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	e := echo.New()

	shortenRepo := repository.NewShortenerRepository(db)
	timeoutContext := time.Duration(config.Timeout) * time.Second
	shortenUse := usecase.NewShortenerUsecase(shortenRepo, timeoutContext, config.Ttl)

	http.NewShortenerHandler(e, shortenUse)

	go func() {
		gocron.Every(uint64(config.CronInterval)).Seconds().Do(shortenRepo.DeleteExpired)
		<-gocron.Start()
	}()

	return e.Start(config.Address)
}

func newDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Hour)
	if err := db2.InitTables(db); err != nil {
		return nil, err
	}

	return db, nil
}
