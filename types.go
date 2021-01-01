package sqlvault

import (
	"context"
	"database/sql"
	sqlDriver "database/sql/driver"
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
	duration    int
	leased      time.Time

	lastConnect time.Time
	flightGroup singleflight.Group
}

type driver interface {
	CreateConnectionString(rawDataSource, username, password string) string
}

// Session is the actual interface exposed as the API
type Session interface {
	// GetConnection returns a valid connection to the database that can be used
	// to make requests to said database.
	// Returns an error if a valid connection could not be obtained
	GetConnection() (*sql.DB, error)
}

// DB is used as an abstraction for the sql.DB struct to allow for better testing
type DB interface {
	Begin() (*sql.Tx, error)
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Close() error
	Conn(context.Context) (*sql.Conn, error)
	Driver() sqlDriver.Driver
	Exec(string, ...interface{}) (sql.Result, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	Ping() error
	PingContext(context.Context) error
	Prepare(string) (*sql.Stmt, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
