package sqlvault

import (
	"database/sql"
	"errors"
)

// ObtainConnection creates a new connection if the old one expired, otherwise does nothing
func (d *db) ObtainConnection() (*sql.DB, error) {
	if d.SQL == nil {
		return nil, errors.New("Has no connection")
	}

	return d.SQL, nil
}
