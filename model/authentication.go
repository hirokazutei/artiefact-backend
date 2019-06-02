package model

// FindValidTokensByUserID finds non-expired tokens with user_id
func FindValidTokensByUserID(tx Queryer, userID int64) ([]AccessToken, error) {
	rows, err := tx.Query(`
		SELECT
			token,
			user_id,
			generated_datetime,
			expiry_datetime,
			obtained_by,
			token_type,
			expired
		FROM
			access_tokens
		WHERE
			user_id = $1
			AND expired = false`,
		userID)
	if err != nil {
		return nil, err
	}
	var ats []AccessToken
	for rows.Next() {
		var at AccessToken
		err := rows.Scan(
			&at.Token,
			&at.UserID,
			&at.GeneratedDatetime,
			&at.ExpiryDatetime,
			&at.ObtainedBy,
			&at.TokenType,
			&at.Expired,
		)
		if err != nil {
			return nil, err
		}
		ats = append(ats, at)
	}
	return ats, nil
}
