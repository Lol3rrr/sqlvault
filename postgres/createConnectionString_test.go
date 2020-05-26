package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConnectionString(t *testing.T) {
	tables := []struct {
		Name               string
		InputRawDataSource string
		InputUsername      string
		InputPassword      string
		Result             string
	}{
		{
			Name:               "Valid Input",
			InputRawDataSource: "sslmode=disable dbname=test",
			InputUsername:      "testUser",
			InputPassword:      "testPassword",
			Result:             "sslmode=disable dbname=test user=testUser password=testPassword",
		},
		{
			Name:               "Still contained Username and Password",
			InputRawDataSource: "password=oldPassword sslmode=disable user=oldUser dbname=test",
			InputUsername:      "testUser",
			InputPassword:      "testPassword",
			Result:             "sslmode=disable dbname=test user=testUser password=testPassword",
		},
	}

	for _, table := range tables {
		inRawDataSource := table.InputRawDataSource
		inUsername := table.InputUsername
		inPassword := table.InputPassword
		result := table.Result

		tmpDriver := &Driver{}

		t.Run(table.Name, func(t *testing.T) {
			output := tmpDriver.CreateConnectionString(inRawDataSource, inUsername, inPassword)

			assert.Equal(t, result, output)
		})
	}
}
