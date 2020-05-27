package sqlvault

import (
	"context"
	"database/sql"
	"time"
)

// ObtainConnection creates a new connection if the old one expired, otherwise does nothing
func (d *DB) ObtainConnection() (*sql.DB, error) {
	d.mux.Lock()
	defer d.mux.Unlock()

	if d.SQL != nil || len(d.username) <= 0 {
		query := "SELECT 1 FROM pg_roles WHERE rolname='" + d.username + "'"

		timeout, cancel := context.WithTimeout(context.TODO(), 100*time.Millisecond)
		defer cancel()
		if _, err := d.SQL.QueryContext(timeout, query); err == nil {
			return d.SQL, nil
		}
	}

	if err := d.connect(); err != nil {
		return nil, err
	}

	return d.SQL, nil
}
