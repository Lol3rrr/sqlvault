package sqlvault

import (
	"errors"
	"fmt"
)

func (db *db) loadCreds() error {
	data, err := db.vaultClient.Logical().Read(db.Settings.VaultCredsPath)
	if err != nil {
		return err
	}
	if data == nil || data.Data == nil {
		return errors.New("Data field was not set in response")
	}

	user, userWorked := data.Data["username"].(string)
	password, passwordWorked := data.Data["password"].(string)

	if !userWorked || !passwordWorked {
		return fmt.Errorf("Data field was malformed: '%+v'", data.Data)
	}

	db.username = user
	db.password = password

	return nil
}
