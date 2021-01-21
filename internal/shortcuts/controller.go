package shortcuts

import (
	"lynx/internal/db"
	"lynx/internal/rest"

	"github.com/kataras/iris/v12"
)

type stringMap map[string]string

type Controller struct {
	factory *Factory
	store   db.URLStore
	vhost   string
}

func NewController(vhost string, store db.URLStore) *Controller {
	return &Controller{
		factory: NewFactory(HashingGenerator),
		store:   store,
		vhost:   vhost,
	}
}

func (c Controller) Shorten(urlString string, shortcode string) *rest.ApiResponse {
	if len(urlString) == 0 {
		return buildApiError(
			iris.StatusLengthRequired,
			"Empty parameter `url`.",
		)
	}

	if !c.factory.isValid(urlString) {
		return buildApiError(
			iris.StatusBadRequest,
			"Not a valid URL.",
		)
	}
	if len(shortcode) <= 0 {
		shortcode, _ = c.factory.gen(urlString)
	}
	if err := c.store.Set(shortcode, urlString); err != nil {
		return buildApiError(
			iris.StatusInternalServerError,
			"Internal error while saving the URL.",
		)
	}

	shortUrl := "http://" + c.vhost + "/u/" + shortcode
	return buildOkResponse(
		rest.DataObject{
			"shortUrl":  shortUrl,
			"shortcode": shortcode,
		},
	)
}

func buildApiError(code int, description string) *rest.ApiResponse {
	return &rest.ApiResponse{
		StatusCode: code,
		Data: rest.DataObject{
			"error": rest.DataObject{
				"description": description,
			},
		},
	}
}

func buildOkResponse(data rest.DataObject) *rest.ApiResponse {
	return &rest.ApiResponse{
		StatusCode: 200,
		Data:       data,
	}
}
