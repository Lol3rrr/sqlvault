package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveUserPassword(t *testing.T) {
	tables := []struct {
		Name         string
		InputString  string
		OutputString string
	}{
		{
			Name:         "Valid Input",
			InputString:  "dbname=test user=test password=yikes sslmode=disable",
			OutputString: "dbname=test sslmode=disable",
		},
	}

	for _, table := range tables {
		inString := table.InputString
		outString := table.OutputString

		t.Run(table.Name, func(t *testing.T) {
			result := removeUserPassword(inString)

			assert.Equal(t, outString, result)
		})
	}
}
