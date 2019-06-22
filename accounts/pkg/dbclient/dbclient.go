package dbclient

type IDbClient interface {
	OpenDb()
	Seed()
}
