package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

// MockSession is used as a simple mock implementation for testing
type MockSession struct {
	mock.Mock
}

// WithRetry is needed to comply with the Session interface
func (m *MockSession) WithRetry(call func(con *sql.DB) error) error {
	args := m.Called(call)
	return args.Error(0)
}
