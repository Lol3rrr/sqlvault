package sqlvault

import "github.com/Lol3rrr/sqlvault/postgres"

var drivers map[string]driver

func init() {
	postgresDriver := &postgres.Driver{}

	drivers = map[string]driver{
		"postgres": postgresDriver,
	}
}
