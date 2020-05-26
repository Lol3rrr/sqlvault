package sqlvault

import "database/sql"

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

	d.SQL = db

	return nil
}
