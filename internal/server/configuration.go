package server

type StoreType int

const DefaultServerPort = "9090"

/*
 * Lynx will use a store to manage shortened URLs.
 * Bolt is good for testing. DynamoDB and Postgres are needed for deployments.
 */
const (
	BoltStore StoreType = iota
	DynamoStore
	PgStore
)

type Configuration struct {
	StoreType StoreType
	Port      string
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		StoreType: BoltStore,
		Port:      DefaultServerPort,
	}
}
