package shortcuts

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortcode(t *testing.T) {
	factory := NewFactory(HashingGenerator)

	code, err := factory.gen(testUrl1)
	assert.Nil(t, err)
	assert.Equal(t, shortcodeLength, len(code))
}

func TestBasicUrlIsValid(t *testing.T) {
	factory := NewFactory(HashingGenerator)
	assert.Equal(t, true, factory.isValid(testUrl1))
}

func TestNonUrlIsNotValid(t *testing.T) {
	factory := NewFactory(HashingGenerator)
	assert.Equal(t, false, factory.isValid("abc"))
}
