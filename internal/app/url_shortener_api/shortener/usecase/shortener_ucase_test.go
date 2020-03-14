package usecase

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"url_shortener_api/internal/app/url_shortener_api/shortener/mocks"
)

func TestEncodeDecode(t *testing.T) {
	var id int64 = 123456
	base62Id := "Gho"
	mockRepo := new(mocks.MockRepository)
	u := NewShortenerUsecase(mockRepo, 2, 10)

	assert.Equal(t, u.Encode(id), base62Id)
	assert.Equal(t, u.Decode(base62Id), id)
}
