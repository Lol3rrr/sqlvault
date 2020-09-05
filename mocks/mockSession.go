package mocks

import (
	"github.com/Lol3rrr/sqlvault"
	"github.com/stretchr/testify/mock"
)

// MockSession is used as a simple mock implementation for testing
// This version does not call the function parameter and therefor does not
// test the actual SQL-Statements and logic in there
type MockSession struct {
	mock.Mock
}

// WithRetry is needed to comply with the Session interface
func (m *MockSession) WithRetry(call func(con sqlvault.DB) error) error {
	args := m.Called(call)
	return args.Error(0)
}
