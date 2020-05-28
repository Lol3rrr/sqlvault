package sqlvault

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
)

// Open is used to create new DB connetion
// dataSourceName contains all the Configuration for the Database, expect username and password
func Open(config Config, vClient *api.Client) (*DB, error) {
	dbDriver, found := drivers[config.DriverName]
	if !found {
		return nil, fmt.Errorf("Could not find driver: '%s'", config.DriverName)
	}

	if config.NewUserThreshold.Nanoseconds() == 0 {
		config.NewUserThreshold = 15 * time.Second
	}

	result := &DB{
		Settings:    config,
		driver:      dbDriver,
		vaultClient: vClient,
	}

	result.Connect()

	return result, nil
}
