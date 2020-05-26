package sqlvault

import (
	"database/sql"

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
}

type driver interface {
	CreateConnectionString(rawDataSource, username, password string) string
}
