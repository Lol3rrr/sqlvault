package sqlvault

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (d *DB) connect() error {
	username, password, err := d.loadCreds()
	if err != nil {
		return err
	}

	dataSource := d.driver.CreateConnectionString(d.dataSourceName, username, password)

	db, err := sql.Open(d.driverName, dataSource)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	if d.SQL != nil {
		if err := d.SQL.Close(); err != nil {
			logrus.Errorf("Could not close previous Connection: 's' \n", err)
		}
	}
	d.SQL = db

	return nil
}
