package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolve(t *testing.T) {
	const foo = "foo"
	var testSet = NewCommandSet()

	t.Run("ByDefault", func(t *testing.T) {
		url := testSet.Resolve(foo)
		assert.Equal(t, GoogleSearchString+foo, url.String())
	})
	t.Run("ByShortcut", func(t *testing.T) {
		url := testSet.Resolve("mail")
		assert.Equal(t, GoogleMailUrlString, url.String())
	})
	t.Run("ByCommand", func(t *testing.T) {
		url := testSet.Resolve("bluejeans " + foo)
		assert.Equal(t, BlueJeansBaseString+foo, url.String())
	})
	t.Run("ByCommandNickname", func(t *testing.T) {
		url := testSet.Resolve("bjn " + foo)
		assert.Equal(t, BlueJeansBaseString+foo, url.String())
	})
	t.Run("ByDetector", func(t *testing.T) {
		var searchText = "skylar"
		url := testSet.Resolve("@" + searchText)
		assert.Equal(t, TwitterBaseString+searchText, url.String())
	})
}

func TestGetCommandInfo(t *testing.T) {
	var testSet = NewCommandSet()

	commands := testSet.GetCommandInfo()
	assert.True(t, len(commands) > 0)
	assert.True(t, len(commands[0].GetName()) > 0)
}

func TestParseSearch(t *testing.T) {
	const input = "bjn 293293333"
	cmd, param := parseSearch(input)

	assert.Equal(t, cmd, "bjn")
	assert.Equal(t, param, "293293333")
}
