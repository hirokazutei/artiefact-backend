package model

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// AlembicVersion represents public.alembic_version
type AlembicVersion struct {
	VersionNum interface{} // version_num
}

// Create inserts the AlembicVersion to the database.
func (r *AlembicVersion) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO alembic_version (version_num) VALUES ($1)`,
		&r.VersionNum)
	if err != nil {
		return errors.Wrap(err, "failed to insert alembic_version")
	}
	return nil
}

// GetAlembicVersionByPk select the AlembicVersion from the database.
func GetAlembicVersionByPk(db Queryer, pk0 interface{}) (*AlembicVersion, error) {
	var r AlembicVersion
	err := db.QueryRow(
		`SELECT version_num FROM alembic_version WHERE version_num = $1`,
		pk0).Scan(&r.VersionNum)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select alembic_version")
	}
	return &r, nil
}

// Profile represents public.profile
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

// Profile_picture represents public.profile picture
type Profile_picture struct {
	UserID    int64          // user_id
	Thumbnail sql.NullString // thumbnail
	Image     sql.NullString // image
}

// Create inserts the Profile_picture to the database.
func (r *Profile_picture) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO profile picture (user_id, thumbnail, image) VALUES ($1, $2, $3)`,
		&r.UserID, &r.Thumbnail, &r.Image)
	if err != nil {
		return errors.Wrap(err, "failed to insert profile picture")
	}
	return nil
}

// GetProfile_pictureByPk select the Profile_picture from the database.
func GetProfile_pictureByPk(db Queryer, pk0 int64) (*Profile_picture, error) {
	var r Profile_picture
	err := db.QueryRow(
		`SELECT user_id, thumbnail, image FROM profile picture WHERE user_id = $1`,
		pk0).Scan(&r.UserID, &r.Thumbnail, &r.Image)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select profile picture")
	}
	return &r, nil
}

// User represents public.user
type User struct {
	ID           int64     // id
	Password     string    // password
	Email        string    // email
	Birthday     time.Time // birthday
	RegisterDate time.Time // register_date
	Status       string    // status
}

// Create inserts the User to the database.
func (r *User) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO user (password, email, birthday, register_date, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		&r.Password, &r.Email, &r.Birthday, &r.RegisterDate, &r.Status).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert user")
	}
	return nil
}

// GetUserByPk select the User from the database.
func GetUserByPk(db Queryer, pk0 int64) (*User, error) {
	var r User
	err := db.QueryRow(
		`SELECT id, password, email, birthday, register_date, status FROM user WHERE id = $1`,
		pk0).Scan(&r.ID, &r.Password, &r.Email, &r.Birthday, &r.RegisterDate, &r.Status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select user")
	}
	return &r, nil
}

// UserAgreement represents public.user_agreement
type UserAgreement struct {
	ID            int64     // id
	UserID        int64     // user_id
	AgreementType string    // agreement_type
	AgreementDate time.Time // agreement_date
}

// Create inserts the UserAgreement to the database.
func (r *UserAgreement) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO user_agreement (user_id, agreement_type, agreement_date) VALUES ($1, $2, $3) RETURNING id`,
		&r.UserID, &r.AgreementType, &r.AgreementDate).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert user_agreement")
	}
	return nil
}

// GetUserAgreementByPk select the UserAgreement from the database.
func GetUserAgreementByPk(db Queryer, pk0 int64) (*UserAgreement, error) {
	var r UserAgreement
	err := db.QueryRow(
		`SELECT id, user_id, agreement_type, agreement_date FROM user_agreement WHERE id = $1`,
		pk0).Scan(&r.ID, &r.UserID, &r.AgreementType, &r.AgreementDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select user_agreement")
	}
	return &r, nil
}

// Username represents public.username
type Username struct {
	UserID      int64  // user_id
	LowerName   string // lower_name
	DisplayName string // display_name
}

// Create inserts the Username to the database.
func (r *Username) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO username (user_id, lower_name, display_name) VALUES ($1, $2, $3)`,
		&r.UserID, &r.LowerName, &r.DisplayName)
	if err != nil {
		return errors.Wrap(err, "failed to insert username")
	}
	return nil
}

// GetUsernameByPk select the Username from the database.
func GetUsernameByPk(db Queryer, pk0 int64) (*Username, error) {
	var r Username
	err := db.QueryRow(
		`SELECT user_id, lower_name, display_name FROM username WHERE user_id = $1`,
		pk0).Scan(&r.UserID, &r.LowerName, &r.DisplayName)
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
