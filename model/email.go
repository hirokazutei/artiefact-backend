package model

import (
	"database/sql"

	"github.com/pkg/errors"
)

// IfEmailExist determines if the email is already taken
func IfEmailExist(db Queryer, email string) (bool, error) {
	var foundEmail string
	err := db.QueryRow(
		`SELECT email FROM artiefact_user WHERE email = $1`,
		email).Scan(&foundEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.Wrap(err, "IfEmailExist failed")
	}
	return true, nil
}
