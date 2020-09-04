package sqlvault

import "database/sql"

// WithRetry executes the given function and provides a connection
// if the function returns an auth error, it automatically retries it
// with a new connection
func (d *db) WithRetry(execFunc func(con *sql.DB) error) error {
	tmpCon, err := d.ObtainConnection()
	if err != nil {
		if tmpCon, err = d.Connect(); err != nil {
			return err
		}
	}

	for i := 0; i < 2; i++ {
		err := execFunc(tmpCon)
		if err != nil {
			if IsAuthError(err) {
				if tmpCon, err = d.Connect(); err != nil {
					return err
				}

				continue
			}

			return err
		}

		return nil
	}

	return nil
}
