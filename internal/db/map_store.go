package db

// In-memory implementtation of URLStore.
// Intended for testing purposes.

type stringMap map[string]string

type MapStore struct {
	entries stringMap
}

// Returns a new BoltStore instance with its connection open.
func NewMapStore() *MapStore {
	return &MapStore{
		entries: make(stringMap),
	}
}

// Set sets a shorten url and its key
// Note: Caller is responsible to generate a key.
func (s *MapStore) Set(key string, value string) error {
	s.entries[key] = value
	return nil
}

// Clear clears all the database entries for the table urls.
func (d *MapStore) Clear() error {
	d.entries = make(stringMap)
	return nil
}

// Get returns a url by its key.
// Returns an empty string if not found.
func (s *MapStore) Get(key string) (value string) {
	value, found := s.entries[key]
	if found {
		return value
	}
	return ""
}

// GetByValue returns all keys for a specific (original) url value.
func (d *MapStore) GetByValue(value string) []string {
	return []string{}
}

// Len returns the number of shortcuts saved
func (s *MapStore) Len() int {
	return len(s.entries)
}

func (s *MapStore) Close() {
}
