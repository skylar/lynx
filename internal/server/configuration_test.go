package server

import(
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
  config := DefaultConfiguration()

  assert.Equal(t, BoltStore, config.StoreType)
  assert.Equal(t, DefaultServerPort, config.Port)
}
