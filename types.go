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

// db represents a single Database instance
type db struct {
	Settings Config

	SQL    *sql.DB
	driver driver

	vaultClient *api.Client
	username    string
	password    string

	lastConnect time.Time
	flightGroup singleflight.Group
}

// Session is the actual interface exposed as the API
type Session interface {
	// WithRetry calls the given function with a valid DB-Connection and returns
	// any actual error (no auth errors) that the function returns
	WithRetry(func(con *sql.DB) error) error
}

type driver interface {
	CreateConnectionString(rawDataSource, username, password string) string
}
