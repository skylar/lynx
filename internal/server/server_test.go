package server

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestServer(t *testing.T) {
	config := DefaultConfiguration()
	server := NewLynxServer(config)

	e := httptest.New(t, server.iris)

	e.GET("/test").Expect().
		Status(httptest.StatusNotFound)
}
