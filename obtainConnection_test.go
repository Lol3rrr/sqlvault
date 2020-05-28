package sqlvault

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObtainConnection(t *testing.T) {
	tables := []struct {
		Name        string
		InputDB     *DB
		ResultError bool
	}{
		{
			Name: "Valid Input",
			InputDB: &DB{
				SQL: &sql.DB{},
			},
			ResultError: false,
		},
		{
			Name: "Connection is not set",
			InputDB: &DB{
				SQL: nil,
			},
			ResultError: true,
		},
	}

	for _, table := range tables {
		inDB := table.InputDB
		resError := table.ResultError

		t.Run(table.Name, func(t *testing.T) {
			_, outErr := inDB.ObtainConnection()

			if resError {
				assert.NotNil(t, outErr)
			} else {
				assert.Nil(t, outErr)
			}
		})
	}
}
