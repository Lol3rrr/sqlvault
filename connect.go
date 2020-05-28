package sqlvault

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

// Connect is used to create a new Connection
func (d *DB) Connect() (*sql.DB, error) {
	// A new connection has been established in the last X seconds
	// return that one
	if !time.Now().After(d.lastConnect.Add(d.Settings.NewUserThreshold)) {
		return d.ObtainConnection()
	}

	// Execute the creation of a new connection only once
	_, err, _ := d.flightGroup.Do("connect", func() (interface{}, error) {
		err := d.loadCreds()
		if err != nil {
			return nil, err
		}

		dataSource := d.driver.CreateConnectionString(d.Settings.DataSourceName, d.username, d.password)

		db, err := sql.Open(d.Settings.DriverName, dataSource)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		if d.SQL != nil {
			if err := d.SQL.Close(); err != nil {
				logrus.Errorf("Could not close previous Connection: '%s' \n", err)
			}
		}

		d.SQL = db
		d.lastConnect = time.Now()

		return db, nil
	})

	if err != nil {
		d.SQL = nil
	}

	return d.SQL, err
}
