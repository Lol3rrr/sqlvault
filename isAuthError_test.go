package sqlvault

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAuthError(t *testing.T) {
	tables := []struct {
		Name       string
		InputError error
		Result     bool
	}{
		{
			Name:       "Permission Error",
			InputError: errors.New("permission denied for testTable"),
			Result:     true,
		},
		{
			Name:       "Syntax Error",
			InputError: errors.New("syntax error at or near"),
			Result:     false,
		},
		{
			Name:       "Password auth error",
			InputError: errors.New("password authentication failed for user"),
			Result:     true,
		},
	}

	for _, table := range tables {
		inError := table.InputError
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := IsAuthError(inError)

			assert.Equal(t, result, output)
		})
	}
}
