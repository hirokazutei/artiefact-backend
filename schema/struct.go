package schema

import "time"

// AccessToken struct for access_token resource
type AccessToken struct {
	Active            bool      `json:"active"`
	ExpiryDatetime    time.Time `json:"expiry_datetime"`
	GeneratedDatetime time.Time `json:"generated_datetime"`
	ObtainedBy        string    `json:"obtained_by"`
	Token             string    `json:"token"`
	UserID            int64     `json:"user_id"`
}

// ArtiefactUser struct for artiefact_user resource
type ArtiefactUser struct {
	Birthday         string    `json:"birthday"`
	ID               int64     `json:"id"`
	Password         string    `json:"password"`
	RegisterDatetime time.Time `json:"register_datetime,omitempty"`
	Status           string    `json:"status,omitempty"`
	Username         string    `json:"username"`
}

// RegisteredEmail struct for registered_email resource
type RegisteredEmail struct {
	Email  string `json:"email"`
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

// ArtiefactUserSignInRequest struct for artiefact_user
// POST: /sign-in
type ArtiefactUserSignInRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// ArtiefactUserSignInResponse struct for artiefact_user
// POST: /sign-in
type ArtiefactUserSignInResponse struct {
	AccessToken   *AccessToken   `json:"access_token,omitempty"`
	ArtiefactUser *ArtiefactUser `json:"artiefact_user,omitempty"`
}

// ArtiefactUserSignUpRequest struct for artiefact_user
// POST: /sign-up
type ArtiefactUserSignUpRequest struct {
	Birthday string `json:"birthday"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// ArtiefactUserSignUpResponse struct for artiefact_user
// POST: /sign-up
type ArtiefactUserSignUpResponse struct {
	AccessToken   *AccessToken   `json:"access_token,omitempty"`
	ArtiefactUser *ArtiefactUser `json:"artiefact_user,omitempty"`
}
