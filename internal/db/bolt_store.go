package db

// Adapted from https://github.com/iris-contrib/examples/url-shortener

import (
	"bytes"

	bolt "go.etcd.io/bbolt"
)

// Panic panics, change it if you don't want to panic on critical INITIALIZE-ONLY-ERRORS
var Panic = func(v interface{}) {
	panic(v)
}

var tableURLs = []byte("urls")

// Bolt implementation of URLStore.
type BoltStore struct {
	db *bolt.DB
}

var _ URLStore = &BoltStore{}

// openDatabase open a new database connection
// and returns its instance.
func openDatabase(stumb string) *bolt.DB {
	// Open the data(base) file in the current working directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(stumb, 0600, nil)
	if err != nil {
		Panic(err)
	}

	// create the buckets here
	tables := [...][]byte{
		tableURLs,
	}

	db.Update(func(tx *bolt.Tx) (err error) {
		for _, table := range tables {
			_, err = tx.CreateBucketIfNotExists(table)
			if err != nil {
				Panic(err)
			}
		}

		return
	})

	return db
}

// Returns a new BoltStore instance with its connection open.
func NewBoltStore(stumb string) *BoltStore {
	return &BoltStore{
		db: openDatabase(stumb),
	}
}

// Set sets a shorten url and its key
// Note: Caller is responsible to generate a key.
func (d *BoltStore) Set(key string, value string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(tableURLs)
		// Generate ID for the url
		// Note: we could use that instead of a random string key
		// but we want to simulate a real-world url shortener
		// so we skip that.
		// id, _ := b.NextSequence()
		if err != nil {
			return err
		}

		k := []byte(key)
		valueB := []byte(value)
		c := b.Cursor()

		found := false
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Equal(valueB, v) {
				found = true
				break
			}
		}
		// if value already exists don't re-put it.
		if found {
			return nil
		}

		return b.Put(k, []byte(value))
	})
}

// Clear clears all the database entries for the table urls.
func (d *BoltStore) Clear() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(tableURLs)
	})
}

// Get returns a url by its key.
//
// Returns an empty string if not found.
func (d *BoltStore) Get(key string) (value string) {
	keyB := []byte(key)
	d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tableURLs)
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Equal(keyB, k) {
				value = string(v)
				break
			}
		}

		return nil
	})

	return
}

// GetByValue returns all keys for a specific (original) url value.
func (d *BoltStore) GetByValue(value string) (keys []string) {
	valueB := []byte(value)
	d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tableURLs)
		if b == nil {
			return nil
		}
		c := b.Cursor()
		// first for the bucket's table "urls"
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Equal(valueB, v) {
				keys = append(keys, string(k))
			}
		}

		return nil
	})

	return
}

// Len returns the number of shortcuts saved
func (d *BoltStore) Len() (num int) {
	d.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(tableURLs)
		if b == nil {
			return nil
		}

		b.ForEach(func([]byte, []byte) error {
			num++
			return nil
		})
		return nil
	})
	return
}

// Close shutdowns the db connection.
func (d *BoltStore) Close() {
	if err := d.db.Close(); err != nil {
		Panic(err)
	}
}
