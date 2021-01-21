package db

// Adapted from https://github.com/iris-contrib/examples/tree/master/url-shortener

// URLStore is the store interface for urls.
type URLStore interface {
	Set(key string, value string) error // error if something went wrong
	Get(key string) string              // empty value if not found
	GetByValue(url string) []string     // empty array if none found
	Len() int                           // should return the number of all the records/tables/buckets
	Clear() error                       // purge all entries
	Close()                             // release the store or ignore
}
