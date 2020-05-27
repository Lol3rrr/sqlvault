package sqlvault

import (
	"database/sql"
	"sync"

	"github.com/hashicorp/vault/api"
)

// DB represents a single Database instance
type DB struct {
	SQL            *sql.DB
	driver         driver
	driverName     string
	dataSourceName string

	vaultClient *api.Client
	credsPath   string
	username    string
	password    string

	mux sync.Mutex
}

type driver interface {
	CreateConnectionString(rawDataSource, username, password string) string
}
