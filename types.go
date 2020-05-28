package sqlvault

import (
	"database/sql"
	"time"

	"github.com/hashicorp/vault/api"
	"golang.org/x/sync/singleflight"
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

	lastConnect time.Time
	flightGroup singleflight.Group
}

type driver interface {
	CreateConnectionString(rawDataSource, username, password string) string
}
