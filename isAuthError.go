package sqlvault

import "strings"

// IsAuthError checks if the error is about authorization
// or insufficient permissions, as they often fall under the same
// group when a user no longer exists
func IsAuthError(err error) bool {
	if strings.Contains(err.Error(), "permission denied") {
		return true
	}

	return false
}
