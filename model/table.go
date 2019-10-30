package model

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// AccessToken represents public.access_token
type AccessToken struct {
	Token             string    // token
	UserID            int64     // user_id
	GeneratedDatetime time.Time // generated_datetime
	ExpiryDatetime    time.Time // expiry_datetime
	ObtainedBy        string    // obtained_by
	Active            bool      // active
}

// Create inserts the AccessToken to the database.
func (r *AccessToken) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO access_token (token, user_id, generated_datetime, expiry_datetime, obtained_by, active) VALUES ($1, $2, $3, $4, $5, $6)`,
		&r.Token, &r.UserID, &r.GeneratedDatetime, &r.ExpiryDatetime, &r.ObtainedBy, &r.Active)
	if err != nil {
		return errors.Wrap(err, "failed to insert access_token")
	}
	return nil
}

// GetAccessTokenByPk select the AccessToken from the database.
func GetAccessTokenByPk(db Queryer, pk0 string) (*AccessToken, error) {
	var r AccessToken
	err := db.QueryRow(
		`SELECT token, user_id, generated_datetime, expiry_datetime, obtained_by, active FROM access_token WHERE token = $1`,
		pk0).Scan(&r.Token, &r.UserID, &r.GeneratedDatetime, &r.ExpiryDatetime, &r.ObtainedBy, &r.Active)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select access_token")
	}
	return &r, nil
}

// Artiefact represents public.artiefact
type Artiefact struct {
	ID        int64          // id
	UserID    sql.NullInt64  // user_id
	CreatedAt time.Time      // created_at
	Hint      sql.NullString // hint
	Type      sql.NullString // type
}

// Create inserts the Artiefact to the database.
func (r *Artiefact) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO artiefact (user_id, created_at, hint, type) VALUES ($1, $2, $3, $4) RETURNING id`,
		&r.UserID, &r.CreatedAt, &r.Hint, &r.Type).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact")
	}
	return nil
}

// GetArtiefactByPk select the Artiefact from the database.
func GetArtiefactByPk(db Queryer, pk0 int64) (*Artiefact, error) {
	var r Artiefact
	err := db.QueryRow(
		`SELECT id, user_id, created_at, hint, type FROM artiefact WHERE id = $1`,
		pk0).Scan(&r.ID, &r.UserID, &r.CreatedAt, &r.Hint, &r.Type)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact")
	}
	return &r, nil
}

// ArtiefactAudio represents public.artiefact_audio
type ArtiefactAudio struct {
	ArtiefactID int64          // artiefact_id
	Description sql.NullString // description
	Duration    float64        // duration
	URI         string         // uri
}

// Create inserts the ArtiefactAudio to the database.
func (r *ArtiefactAudio) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO artiefact_audio (artiefact_id, description, duration, uri) VALUES ($1, $2, $3, $4)`,
		&r.ArtiefactID, &r.Description, &r.Duration, &r.URI)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_audio")
	}
	return nil
}

// GetArtiefactAudioByPk select the ArtiefactAudio from the database.
func GetArtiefactAudioByPk(db Queryer, pk0 int64) (*ArtiefactAudio, error) {
	var r ArtiefactAudio
	err := db.QueryRow(
		`SELECT artiefact_id, description, duration, uri FROM artiefact_audio WHERE artiefact_id = $1`,
		pk0).Scan(&r.ArtiefactID, &r.Description, &r.Duration, &r.URI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_audio")
	}
	return &r, nil
}

// ArtiefactDiscovery represents public.artiefact_discovery
type ArtiefactDiscovery struct {
	ArtiefactID int64         // artiefact_id
	UserID      sql.NullInt64 // user_id
	UploadedAt  time.Time     // uploaded_at
}

// Create inserts the ArtiefactDiscovery to the database.
func (r *ArtiefactDiscovery) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO artiefact_discovery (artiefact_id, user_id, uploaded_at) VALUES ($1, $2, $3)`,
		&r.ArtiefactID, &r.UserID, &r.UploadedAt)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_discovery")
	}
	return nil
}

// GetArtiefactDiscoveryByPk select the ArtiefactDiscovery from the database.
func GetArtiefactDiscoveryByPk(db Queryer, pk0 int64) (*ArtiefactDiscovery, error) {
	var r ArtiefactDiscovery
	err := db.QueryRow(
		`SELECT artiefact_id, user_id, uploaded_at FROM artiefact_discovery WHERE artiefact_id = $1`,
		pk0).Scan(&r.ArtiefactID, &r.UserID, &r.UploadedAt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_discovery")
	}
	return &r, nil
}

// ArtiefactImage represents public.artiefact_image
type ArtiefactImage struct {
	ArtiefactID int64          // artiefact_id
	Description sql.NullString // description
	URI         string         // uri
}

// Create inserts the ArtiefactImage to the database.
func (r *ArtiefactImage) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO artiefact_image (artiefact_id, description, uri) VALUES ($1, $2, $3)`,
		&r.ArtiefactID, &r.Description, &r.URI)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_image")
	}
	return nil
}

// GetArtiefactImageByPk select the ArtiefactImage from the database.
func GetArtiefactImageByPk(db Queryer, pk0 int64) (*ArtiefactImage, error) {
	var r ArtiefactImage
	err := db.QueryRow(
		`SELECT artiefact_id, description, uri FROM artiefact_image WHERE artiefact_id = $1`,
		pk0).Scan(&r.ArtiefactID, &r.Description, &r.URI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_image")
	}
	return &r, nil
}

// ArtiefactLocation represents public.artiefact_location
type ArtiefactLocation struct {
	ArtiefactID int64         // artiefact_id
	Longitude   sql.NullInt64 // longitude
	Latitude    sql.NullInt64 // latitude
}

// Create inserts the ArtiefactLocation to the database.
func (r *ArtiefactLocation) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO artiefact_location (artiefact_id, longitude, latitude) VALUES ($1, $2, $3)`,
		&r.ArtiefactID, &r.Longitude, &r.Latitude)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_location")
	}
	return nil
}

// GetArtiefactLocationByPk select the ArtiefactLocation from the database.
func GetArtiefactLocationByPk(db Queryer, pk0 int64) (*ArtiefactLocation, error) {
	var r ArtiefactLocation
	err := db.QueryRow(
		`SELECT artiefact_id, longitude, latitude FROM artiefact_location WHERE artiefact_id = $1`,
		pk0).Scan(&r.ArtiefactID, &r.Longitude, &r.Latitude)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_location")
	}
	return &r, nil
}

// ArtiefactRating represents public.artiefact_rating
type ArtiefactRating struct {
	ArtiefactID int64         // artiefact_id
	UserID      sql.NullInt64 // user_id
	RatedAt     time.Time     // rated_at
	Rating      int16         // rating
}

// Create inserts the ArtiefactRating to the database.
func (r *ArtiefactRating) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO artiefact_rating (artiefact_id, user_id, rated_at, rating) VALUES ($1, $2, $3, $4)`,
		&r.ArtiefactID, &r.UserID, &r.RatedAt, &r.Rating)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_rating")
	}
	return nil
}

// GetArtiefactRatingByPk select the ArtiefactRating from the database.
func GetArtiefactRatingByPk(db Queryer, pk0 int64) (*ArtiefactRating, error) {
	var r ArtiefactRating
	err := db.QueryRow(
		`SELECT artiefact_id, user_id, rated_at, rating FROM artiefact_rating WHERE artiefact_id = $1`,
		pk0).Scan(&r.ArtiefactID, &r.UserID, &r.RatedAt, &r.Rating)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_rating")
	}
	return &r, nil
}

// ArtiefactText represents public.artiefact_text
type ArtiefactText struct {
	ArtiefactID int64          // artiefact_id
	Title       string         // title
	Text        sql.NullString // text
}

// Create inserts the ArtiefactText to the database.
func (r *ArtiefactText) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO artiefact_text (artiefact_id, title, text) VALUES ($1, $2, $3)`,
		&r.ArtiefactID, &r.Title, &r.Text)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_text")
	}
	return nil
}

// GetArtiefactTextByPk select the ArtiefactText from the database.
func GetArtiefactTextByPk(db Queryer, pk0 int64) (*ArtiefactText, error) {
	var r ArtiefactText
	err := db.QueryRow(
		`SELECT artiefact_id, title, text FROM artiefact_text WHERE artiefact_id = $1`,
		pk0).Scan(&r.ArtiefactID, &r.Title, &r.Text)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_text")
	}
	return &r, nil
}

// ArtiefactUser represents public.artiefact_user
type ArtiefactUser struct {
	ID               int64     // id
	Password         string    // password
	Birthday         time.Time // birthday
	RegisterDatetime time.Time // register_datetime
	Status           string    // status
}

// Create inserts the ArtiefactUser to the database.
func (r *ArtiefactUser) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO artiefact_user (password, birthday, register_datetime, status) VALUES ($1, $2, $3, $4) RETURNING id`,
		&r.Password, &r.Birthday, &r.RegisterDatetime, &r.Status).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_user")
	}
	return nil
}

// GetArtiefactUserByPk select the ArtiefactUser from the database.
func GetArtiefactUserByPk(db Queryer, pk0 int64) (*ArtiefactUser, error) {
	var r ArtiefactUser
	err := db.QueryRow(
		`SELECT id, password, birthday, register_datetime, status FROM artiefact_user WHERE id = $1`,
		pk0).Scan(&r.ID, &r.Password, &r.Birthday, &r.RegisterDatetime, &r.Status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_user")
	}
	return &r, nil
}

// ArtiefactVideo represents public.artiefact_video
type ArtiefactVideo struct {
	ArtiefactID int64          // artiefact_id
	Description sql.NullString // description
	Duration    float64        // duration
	URI         string         // uri
}

// Create inserts the ArtiefactVideo to the database.
func (r *ArtiefactVideo) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO artiefact_video (artiefact_id, description, duration, uri) VALUES ($1, $2, $3, $4)`,
		&r.ArtiefactID, &r.Description, &r.Duration, &r.URI)
	if err != nil {
		return errors.Wrap(err, "failed to insert artiefact_video")
	}
	return nil
}

// GetArtiefactVideoByPk select the ArtiefactVideo from the database.
func GetArtiefactVideoByPk(db Queryer, pk0 int64) (*ArtiefactVideo, error) {
	var r ArtiefactVideo
	err := db.QueryRow(
		`SELECT artiefact_id, description, duration, uri FROM artiefact_video WHERE artiefact_id = $1`,
		pk0).Scan(&r.ArtiefactID, &r.Description, &r.Duration, &r.URI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select artiefact_video")
	}
	return &r, nil
}

// EmailVerification represents public.email_verification
type EmailVerification struct {
	RequestID            int64     // request_id
	VerificationDatetime time.Time // verification_datetime
}

// Create inserts the EmailVerification to the database.
func (r *EmailVerification) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO email_verification (request_id, verification_datetime) VALUES ($1, $2)`,
		&r.RequestID, &r.VerificationDatetime)
	if err != nil {
		return errors.Wrap(err, "failed to insert email_verification")
	}
	return nil
}

// GetEmailVerificationByPk select the EmailVerification from the database.
func GetEmailVerificationByPk(db Queryer, pk0 int64) (*EmailVerification, error) {
	var r EmailVerification
	err := db.QueryRow(
		`SELECT request_id, verification_datetime FROM email_verification WHERE request_id = $1`,
		pk0).Scan(&r.RequestID, &r.VerificationDatetime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select email_verification")
	}
	return &r, nil
}

// EmailVerificationRequest represents public.email_verification_request
type EmailVerificationRequest struct {
	RegisteredEmailID  int64     // registered_email_id
	Code               string    // code
	ExpirationDatetime time.Time // expiration_datetime
	RequestedDatetime  time.Time // requested_datetime
}

// Create inserts the EmailVerificationRequest to the database.
func (r *EmailVerificationRequest) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO email_verification_request (registered_email_id, code, expiration_datetime, requested_datetime) VALUES ($1, $2, $3, $4)`,
		&r.RegisteredEmailID, &r.Code, &r.ExpirationDatetime, &r.RequestedDatetime)
	if err != nil {
		return errors.Wrap(err, "failed to insert email_verification_request")
	}
	return nil
}

// GetEmailVerificationRequestByPk select the EmailVerificationRequest from the database.
func GetEmailVerificationRequestByPk(db Queryer, pk0 int64) (*EmailVerificationRequest, error) {
	var r EmailVerificationRequest
	err := db.QueryRow(
		`SELECT registered_email_id, code, expiration_datetime, requested_datetime FROM email_verification_request WHERE registered_email_id = $1`,
		pk0).Scan(&r.RegisteredEmailID, &r.Code, &r.ExpirationDatetime, &r.RequestedDatetime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select email_verification_request")
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

// ProfilePicture represents public.profile_picture
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

// RegisteredEmail represents public.registered_email
type RegisteredEmail struct {
	ID         int64  // id
	UserID     int64  // user_id
	Email      string // email
	EmailLower string // email_lower
	Status     string // status
}

// Create inserts the RegisteredEmail to the database.
func (r *RegisteredEmail) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO registered_email (user_id, email, email_lower, status) VALUES ($1, $2, $3, $4) RETURNING id`,
		&r.UserID, &r.Email, &r.EmailLower, &r.Status).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert registered_email")
	}
	return nil
}

// GetRegisteredEmailByPk select the RegisteredEmail from the database.
func GetRegisteredEmailByPk(db Queryer, pk0 int64) (*RegisteredEmail, error) {
	var r RegisteredEmail
	err := db.QueryRow(
		`SELECT id, user_id, email, email_lower, status FROM registered_email WHERE id = $1`,
		pk0).Scan(&r.ID, &r.UserID, &r.Email, &r.EmailLower, &r.Status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select registered_email")
	}
	return &r, nil
}

// TokenAccess represents public.token_access
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

// UserAgreement represents public.user_agreement
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

// Username represents public.username
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
