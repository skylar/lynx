package server

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestServer(t *testing.T) {
	testUrl := "http://somewhere.over.rainbows/"
	shortcode := "oz"
	config := DefaultConfiguration()
	server := NewLynxServer(config)

	e := httptest.New(t, server.iris)

	e.GET("/test").Expect().
		Status(httptest.StatusNotFound)

	e.POST("/api/shorten").
		WithFormField("url", testUrl).
		WithFormField("shortcode", shortcode).Expect().
		Status(httptest.StatusOK).Body().Contains("{\"shortUrl\":")

	e.GET("/u/" + shortcode).Expect().
		Status(httptest.StatusTemporaryRedirect).Header("Location").Equal(testUrl)
}
