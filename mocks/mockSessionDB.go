package mocks

import (
	"github.com/Lol3rrr/sqlvault"
)

// MockSessionDB is used as a simple mock implementation for testing
// This version does call the function parameter
type MockSessionDB struct {
	mockDB *MockDB
}

// WithRetry is needed to comply with the interface
func (m *MockSessionDB) WithRetry(call func(con sqlvault.DB) error) error {
	return call(m.mockDB)
}
