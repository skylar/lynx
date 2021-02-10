package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const testKey = "key";
const testValue = "value";

func TestSetGetClear(t *testing.T) {
	store := NewMapStore();

	err := store.Set(testKey, testValue);
	assert.Nil(t, err)
	assert.Equal(t, 1, store.Len())

	value := store.Get(testKey)
	assert.Equal(t, testValue, value)

	err = store.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, store.Len())
}
