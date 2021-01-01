package sqlvault

import (
	"database/sql"
	"time"
)

func (d *db) GetConnection() (*sql.DB, error) {
	if d.SQL == nil || int(time.Since(d.leased).Seconds()) > d.duration {
		return d.Connect()
	}

	return d.SQL, nil
}
