package model

import (
	"database/sql"
	"strings"

	"github.com/pkg/errors"
)

// IfUsernameExist determines if the Usermane is already taken
func IfUsernameExist(db Queryer, username string) (bool, error) {
	var foundUsername string
	err := db.QueryRow(
		`SELECT
			username_lower
		FROM
			username
		WHERE
			username_lower = $1`,
		strings.ToLower(username)).Scan(&foundUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.Wrap(err, "IfUsernameExist failed")
	}
	return true, nil
}

// GetActiveArtiefactUserByID obtains ArtiefactUser by ID
func GetActiveArtiefactUserByID(db Queryer, id int64) (*ArtiefactUser, bool, error) {
	var au ArtiefactUser
	err := db.QueryRow(
		`SELECT
			id,
			password,
			birthday
			register_datetime
			status
		FROM
			artiefact_user
		WHERE
			id = $1`, id).Scan(
		&au.ID,
		&au.Password,
		&au.Birthday,
		&au.RegisterDatetime,
		&au.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, errors.Wrap(err, "IfUsernameExist failed")
	}
	return &au, true, nil
}

// GetArtiefactUserByUsername obtains ArtiefactUser by Username
func GetArtiefactUserByUsername(db Queryer, username string) (*ArtiefactUser, error) {
	var au ArtiefactUser
	err := db.QueryRow(
		`SELECT
			id,
			password,
			birthday
			register_datetime
			status
		FROM
			artiefact_user au
		JOIN
			username u
		ON
			u.user_id = au.id
		WHERE
			username_lower = $1`,
		strings.ToLower(username)).Scan(
		&au.ID,
		&au.Password,
		&au.Birthday,
		&au.RegisterDatetime,
		&au.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "IfUsernameExist failed")
	}
	return &au, nil
}
