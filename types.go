package sqlvault

import (
	"database/sql"
	"time"

	"github.com/hashicorp/vault/api"
	"golang.org/x/sync/singleflight"
)

// Config contains all the Options that can be used to
// configure the behaviour of this package
type Config struct {
	DriverName     string
	DataSourceName string

	VaultCredsPath string

	// Defaults to 15 Seconds if not defined
	NewUserThreshold time.Duration
}

// DB represents a single Database instance
type DB struct {
	Settings Config

	SQL    *sql.DB
	driver driver

	vaultClient *api.Client
	username    string
	password    string

	lastConnect time.Time
	flightGroup singleflight.Group
}

type driver interface {
	CreateConnectionString(rawDataSource, username, password string) string
}
