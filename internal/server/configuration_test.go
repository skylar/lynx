package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	config := DefaultConfiguration()

	assert.Equal(t, BoltStoreType, config.StoreType)
	assert.Equal(t, DefaultStoreConnectionString, config.StoreConnectionString)
	assert.Equal(t, DefaultServerPort, config.Port)
}
