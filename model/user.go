package model

import (
	"database/sql"
	"strings"

	"github.com/pkg/errors"
)

// CheckIfUsernameIsTaken determines if the Usermane is already taken
func CheckIfUsernameIsTaken(db Queryer, username string) (bool, error) {
	var foundUsername string
	err := db.QueryRow(
		`SELECT username_lower FROM username WHERE username_lower = $1`,
		strings.ToLower(username)).Scan(&foundUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.Wrap(err, "CheckIfUsernameIsTaken failed")
	}
	return true, nil
}
