package model

import (
	"database/sql"

	"github.com/pkg/errors"
)

// GetValidToken finds non-expired tokens
func GetValidToken(db Queryer, token string) (*AccessToken, bool, error) {
	var t AccessToken
	err := db.QueryRow(`
		SELECT
			token,
			user_id,
			generated_datetime,
			expiry_datetime,
			obtained_by,
			active
		FROM
			access_token
		WHERE
			token = $1
			AND active = true`,
		token).Scan(
		&t.Token,
		&t.UserID,
		&t.GeneratedDatetime,
		&t.ExpiryDatetime,
		&t.ObtainedBy,
		&t.Active,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, errors.Wrap(err, "GetValidToken failed")
	}

	return &t, true, nil
}
