package model

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// AccessToken represents artiefact.access_token
type AccessToken struct {
	Token             string    // token
	UserID            int64     // user_id
	GeneratedDatetime time.Time // generated_datetime
	ExpiryDatetime    time.Time // expiry_datetime
	ObtainedBy        string    // obtained_by
	TokenType         string    // token_type
}

// Create inserts the AccessToken to the database.
func (r *AccessToken) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO access_token (token, user_id, generated_datetime, expiry_datetime, obtained_by, token_type) VALUES ($1, $2, $3, $4, $5, $6)`,
		&r.Token, &r.UserID, &r.GeneratedDatetime, &r.ExpiryDatetime, &r.ObtainedBy, &r.TokenType)
	if err != nil {
		return errors.Wrap(err, "failed to insert access_token")
	}
	return nil
}

// GetAccessTokenByPk select the AccessToken from the database.
func GetAccessTokenByPk(db Queryer, pk0 string) (*AccessToken, error) {
	var r AccessToken
	err := db.QueryRow(
		`SELECT token, user_id, generated_datetime, expiry_datetime, obtained_by, token_type FROM access_token WHERE token = $1`,
		pk0).Scan(&r.Token, &r.UserID, &r.GeneratedDatetime, &r.ExpiryDatetime, &r.ObtainedBy, &r.TokenType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select access_token")
	}
	return &r, nil
}

// ArtiefactUser represents artiefact.artiefact_user
type ArtiefactUser struct {
	ID               int64     // id
	Password         string    // password
	Email            string    // email
	Birthday         time.Time // birthday
	RegisterDatetime time.Time // register_datetime
	Status           string    // status
}

// Create inserts the ArtiefactUser to the database.
func (r *ArtiefactUser) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO artiefact_user (password, email, birthday, register_datetime, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		&r.Password, &r.Email, &r.Birthday, &r.RegisterDatetime, &r.Status).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_user")
	}
	return nil
}

// GetArtiefactUserByPk select the ArtiefactUser from the database.
func GetArtiefactUserByPk(db Queryer, pk0 int64) (*ArtiefactUser, error) {
	var r ArtiefactUser
	err := db.QueryRow(
		`SELECT id, password, email, birthday, register_datetime, status FROM artiefact_user WHERE id = $1`,
		pk0).Scan(&r.ID, &r.Password, &r.Email, &r.Birthday, &r.RegisterDatetime, &r.Status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_user")
	}
	return &r, nil
}

// Profile represents artiefact.profile
type Profile struct {
	UserID  int64          // user_id
	Name    sql.NullString // name
	Website sql.NullString // website
	Bio     sql.NullString // bio
	Gender  sql.NullInt64  // gender
}

// Create inserts the Profile to the database.
func (r *Profile) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO profile (user_id, name, website, bio, gender) VALUES ($1, $2, $3, $4, $5)`,
		&r.UserID, &r.Name, &r.Website, &r.Bio, &r.Gender)
	if err != nil {
		return errors.Wrap(err, "failed to insert profile")
	}
	return nil
}

// GetProfileByPk select the Profile from the database.
func GetProfileByPk(db Queryer, pk0 int64) (*Profile, error) {
	var r Profile
	err := db.QueryRow(
		`SELECT user_id, name, website, bio, gender FROM profile WHERE user_id = $1`,
		pk0).Scan(&r.UserID, &r.Name, &r.Website, &r.Bio, &r.Gender)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select profile")
	}
	return &r, nil
}

// ProfilePicture represents artiefact.profile_picture
type ProfilePicture struct {
	UserID    int64          // user_id
	Thumbnail sql.NullString // thumbnail
	Image     sql.NullString // image
}

// Create inserts the ProfilePicture to the database.
func (r *ProfilePicture) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO profile_picture (user_id, thumbnail, image) VALUES ($1, $2, $3)`,
		&r.UserID, &r.Thumbnail, &r.Image)
	if err != nil {
		return errors.Wrap(err, "failed to insert profile_picture")
	}
	return nil
}

// GetProfilePictureByPk select the ProfilePicture from the database.
func GetProfilePictureByPk(db Queryer, pk0 int64) (*ProfilePicture, error) {
	var r ProfilePicture
	err := db.QueryRow(
		`SELECT user_id, thumbnail, image FROM profile_picture WHERE user_id = $1`,
		pk0).Scan(&r.UserID, &r.Thumbnail, &r.Image)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select profile_picture")
	}
	return &r, nil
}

// TokenAccess represents artiefact.token_access
type TokenAccess struct {
	ID               int64     // id
	Token            string    // token
	LastUsedDatetime time.Time // last_used_datetime
}

// Create inserts the TokenAccess to the database.
func (r *TokenAccess) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO token_access (token, last_used_datetime) VALUES ($1, $2) RETURNING id`,
		&r.Token, &r.LastUsedDatetime).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert token_access")
	}
	return nil
}

// GetTokenAccessByPk select the TokenAccess from the database.
func GetTokenAccessByPk(db Queryer, pk0 int64) (*TokenAccess, error) {
	var r TokenAccess
	err := db.QueryRow(
		`SELECT id, token, last_used_datetime FROM token_access WHERE id = $1`,
		pk0).Scan(&r.ID, &r.Token, &r.LastUsedDatetime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select token_access")
	}
	return &r, nil
}

// UserAgreement represents artiefact.user_agreement
type UserAgreement struct {
	ID                int64     // id
	UserID            int64     // user_id
	AgreementType     string    // agreement_type
	AgreementDatetime time.Time // agreement_datetime
}

// Create inserts the UserAgreement to the database.
func (r *UserAgreement) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO user_agreement (user_id, agreement_type, agreement_datetime) VALUES ($1, $2, $3) RETURNING id`,
		&r.UserID, &r.AgreementType, &r.AgreementDatetime).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert user_agreement")
	}
	return nil
}

// GetUserAgreementByPk select the UserAgreement from the database.
func GetUserAgreementByPk(db Queryer, pk0 int64) (*UserAgreement, error) {
	var r UserAgreement
	err := db.QueryRow(
		`SELECT id, user_id, agreement_type, agreement_datetime FROM user_agreement WHERE id = $1`,
		pk0).Scan(&r.ID, &r.UserID, &r.AgreementType, &r.AgreementDatetime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select user_agreement")
	}
	return &r, nil
}

// Username represents artiefact.username
type Username struct {
	UserID        int64  // user_id
	UsernameLower string // username_lower
	UsernameRaw   string // username_raw
}

// Create inserts the Username to the database.
func (r *Username) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO username (user_id, username_lower, username_raw) VALUES ($1, $2, $3)`,
		&r.UserID, &r.UsernameLower, &r.UsernameRaw)
	if err != nil {
		return errors.Wrap(err, "failed to insert username")
	}
	return nil
}

// GetUsernameByPk select the Username from the database.
func GetUsernameByPk(db Queryer, pk0 int64) (*Username, error) {
	var r Username
	err := db.QueryRow(
		`SELECT user_id, username_lower, username_raw FROM username WHERE user_id = $1`,
		pk0).Scan(&r.UserID, &r.UsernameLower, &r.UsernameRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select username")
	}
	return &r, nil
}

// Queryer database/sql compatible query interface
type Queryer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}
