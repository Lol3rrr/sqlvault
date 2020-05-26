package sqlvault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

// Open is used to create new DB connetion
// dataSourceName contains all the Configuration for the Database, expect username and password
func Open(driverName, dataSourceName string, vClient *api.Client, dbCredsPath string) (*DB, error) {
	dbDriver, found := drivers[driverName]
	if !found {
		return nil, fmt.Errorf("Could not find driver: '%s'", driverName)
	}

	result := &DB{
		driver:         dbDriver,
		driverName:     driverName,
		dataSourceName: dataSourceName,

		vaultClient: vClient,
		credsPath:   dbCredsPath,
	}

	result.connect()

	return result, nil
}
