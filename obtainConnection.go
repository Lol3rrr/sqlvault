package sqlvault

import (
	"context"
	"database/sql"
	"time"
)

// ObtainConnection creates a new connection if the old one expired, otherwise does nothing
func (d *DB) ObtainConnection(tableName string) (*sql.DB, error) {
	if d.SQL != nil {
		query := "SELECT * FROM " + tableName + " WHERE 0;"

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
