package sqlvault

import "database/sql"

// ObtainConnection creates a new connection if the old one expired, otherwise does nothing
func (d *DB) ObtainConnection() (*sql.DB, error) {
	err := d.SQL.Ping()
	if err == nil {
		return d.SQL, nil
	}

	if err := d.connect(); err != nil {
		return nil, err
	}

	return d.SQL, nil
}
