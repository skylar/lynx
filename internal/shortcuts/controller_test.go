package shortcuts

import (
	"lynx/internal/db"
	"lynx/internal/rest"

	"testing"
	"github.com/kataras/iris/v12"
	"github.com/stretchr/testify/assert"
)

const testUrl1 = "https://www.reddit.com/"
const testUrl2 = "https://www.discord.com/"
const emptyString = ""

func makeTestController() *Controller {
	return NewController("shorten.server.test", db.NewMapStore())
}

func assertValidShortcodeResponse(t *testing.T, response *rest.ApiResponse) string {
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, response.Data["shortUrl"])
	return response.Data["shortcode"].(string);
}

func TestShorten(t *testing.T) {
	controller := makeTestController()

	response := controller.Shorten(testUrl1, emptyString)
	shortcode := assertValidShortcodeResponse(t, response)
	assert.Equal(t, shortcodeLength, len(shortcode))

	similarResponse := controller.Shorten(testUrl1, emptyString)
	assert.Equal(t, shortcode, assertValidShortcodeResponse(t, similarResponse))
	assert.Equal(t, 1, controller.store.Len())

	differentResponse := controller.Shorten(testUrl2, emptyString)
	assert.NotEqual(t, shortcode, assertValidShortcodeResponse(t, differentResponse))
	assert.Equal(t, 2, controller.store.Len())
}

func TestShortenWithShortcode(t *testing.T) {
	testShortcode := "reddit"
	controller := makeTestController()

	response := controller.Shorten(testUrl1, testShortcode)
	shortcode := assertValidShortcodeResponse(t, response)
	assert.Equal(t, testShortcode, shortcode)
}

func TestShortenInvalidUrlReturnsError(t *testing.T) {
	controller := makeTestController()

	response := controller.Shorten("abc", "")
	assert.Equal(t, iris.StatusBadRequest, response.StatusCode)
}
