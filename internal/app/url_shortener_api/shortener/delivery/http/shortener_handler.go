package http

import (
	"context"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"url_shortener_api/internal/app/url_shortener_api/models"
	"url_shortener_api/internal/app/url_shortener_api/shortener"
)
import "github.com/labstack/echo/v4"

type ResponseError struct {
	Message string `json:"message"`
}

type ShortenerHandler struct {
	SUsecase shortener.Usecase
}

func NewShortenerHandler(e *echo.Echo, us shortener.Usecase) {
	handler := &ShortenerHandler{
		SUsecase: us,
	}

	e.GET("/new_short_url", handler.NewShortUrl)
	e.GET("/:short_link", handler.Redirect)
	e.POST("/admin/set_ttl ", handler.SetTtl)
	e.GET("/admin/get_all", handler.GetAll)
}

func (h *ShortenerHandler) NewShortUrl(c echo.Context) error {
	url := c.QueryParam("url")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	shortUrl, err := h.SUsecase.NewShortLink(ctx, url)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, shortUrl)
}

func (h *ShortenerHandler) Redirect(c echo.Context) error {
	shortLink := c.Param("short_link")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	url, err := h.SUsecase.Get(ctx, shortLink)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, url)
	//return c.JSON(http.StatusOK, url)
}

func (h *ShortenerHandler) SetTtl(c echo.Context) error {
	var ttlObj models.Ttl
	err := c.Bind(&ttlObj)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&ttlObj); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = h.SUsecase.SetTtl(ttlObj.Ttl)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "Hello world. It is URL Shortener")
}

func (h *ShortenerHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	links, err := h.SUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, links)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	fmt.Println(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *models.Ttl) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
