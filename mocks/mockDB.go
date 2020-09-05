package mocks

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/stretchr/testify/mock"
)

// MockDB is used as a mock implementation for the DB interface
// for easier testing
type MockDB struct {
	mock.Mock
}

// Begin is needed to comply with the interface
func (m *MockDB) Begin() (*sql.Tx, error) {
	args := m.Called()
	return args.Get(0).(*sql.Tx), args.Error(1)
}

// BeginTx is needed to comply with the interface
func (m *MockDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*sql.Tx), args.Error(1)
}

// Close is needed to comply with the interface
func (m *MockDB) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Conn is needed to comply with the interface
func (m *MockDB) Conn(ctx context.Context) (*sql.Conn, error) {
	args := m.Called(ctx)
	return args.Get(0).(*sql.Conn), args.Error(1)
}

// Driver is needed to comply with the interface
func (m *MockDB) Driver() driver.Driver {
	args := m.Called()
	return args.Get(0).(driver.Driver)
}

// Exec is needed to comply with the interface
func (m *MockDB) Exec(query string, queryArgs ...interface{}) (sql.Result, error) {
	args := m.Called(query, queryArgs)
	return args.Get(0).(sql.Result), args.Error(1)
}

// ExecContext is needed to comply with the interface
func (m *MockDB) ExecContext(ctx context.Context, query string, queryArgs ...interface{}) (sql.Result, error) {
	args := m.Called(ctx, query, queryArgs)
	return args.Get(0).(sql.Result), args.Error(1)
}

// Ping is needed to comply with the interface
func (m *MockDB) Ping() error {
	args := m.Called()
	return args.Error(0)
}

// PingContext is needed to comply with the interface
func (m *MockDB) PingContext(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Prepare is needed to comply with the interface
func (m *MockDB) Prepare(query string) (*sql.Stmt, error) {
	args := m.Called(query)
	return args.Get(0).(*sql.Stmt), args.Error(1)
}

// PrepareContext is needed to comply with the interface
func (m *MockDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	args := m.Called(ctx, query)
	return args.Get(0).(*sql.Stmt), args.Error(1)
}

// Query is needed to comply with the interface
func (m *MockDB) Query(query string, queryArgs ...interface{}) (*sql.Rows, error) {
	args := m.Called(query, queryArgs)
	return args.Get(0).(*sql.Rows), args.Error(1)
}

// QueryContext is needed to comply with the interface
func (m *MockDB) QueryContext(ctx context.Context, query string, queryArgs ...interface{}) (*sql.Rows, error) {
	args := m.Called(ctx, query, queryArgs)
	return args.Get(0).(*sql.Rows), args.Error(1)
}

// QueryRow is needed to comply with the interface
func (m *MockDB) QueryRow(query string, queryArgs ...interface{}) *sql.Row {
	args := m.Called(query, queryArgs)
	return args.Get(0).(*sql.Row)
}

// QueryRowContext is needed to comply with the interface
func (m *MockDB) QueryRowContext(ctx context.Context, query string, queryArgs ...interface{}) *sql.Row {
	args := m.Called(ctx, query, queryArgs)
	return args.Get(0).(*sql.Row)
}
