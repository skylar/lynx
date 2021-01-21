package db

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBoltStore_SetGetClear(t *testing.T) {
	testDbName := "bolt-test.db"
	store := NewBoltStore(testDbName)

	err := store.Set(testKey, testValue)
	assert.Nil(t, err)
	assert.Equal(t, 1, store.Len())

	value := store.Get(testKey)
	assert.Equal(t, testValue, value)

	err = store.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, store.Len())

	os.Remove(testDbName)
}
