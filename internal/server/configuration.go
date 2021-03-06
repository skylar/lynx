package server

type StoreType int

const DefaultServerPort = "9090"
const DefaultStoreConnectionString = "none"

/*
 * Lynx will use a store to manage shortened URLs.
 * Bolt is good for testing. DynamoDB and Postgres are needed for deployments.
 */
const (
	BoltStoreType StoreType = iota
	DynamoStoretype
	MapStoreType
	PgStoreType
)

type Configuration struct {
	StoreType             StoreType
	StoreConnectionString string
	Port                  string
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		StoreType:             MapStoreType,
		StoreConnectionString: DefaultStoreConnectionString,
		Port:                  DefaultServerPort,
	}
}
