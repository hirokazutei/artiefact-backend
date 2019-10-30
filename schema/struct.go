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

// ArtiefactObject struct for artiefact_object resource
type ArtiefactObject struct {
	Hint      string  `json:"hint"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Text      string  `json:"text,omitempty"`
	Title     string  `json:"title,omitempty"`
	Type      string  `json:"type"`
	UserID    int64   `json:"user_id"`
}

// ArtiefactUser struct for artiefact_user resource
type ArtiefactUser struct {
	Birthday         string    `json:"birthday"`
	ID               int64     `json:"id"`
	Password         string    `json:"password,omitempty"`
	RegisterDatetime time.Time `json:"register_datetime,omitempty"`
	Status           string    `json:"status"`
	Username         string    `json:"username"`
}

// RegisteredEmail struct for registered_email resource
type RegisteredEmail struct {
	Email  string `json:"email"`
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

// ArtiefactObjectPostTextRequest struct for artiefact_object
// POST: /artiefact/text
type ArtiefactObjectPostTextRequest struct {
	Hint      string  `json:"hint"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Text      string  `json:"text,omitempty"`
	Title     string  `json:"title,omitempty"`
	Type      string  `json:"type"`
	UserID    int64   `json:"user_id"`
}

// ArtiefactObjectPostTextResponse struct for artiefact_object
// POST: /artiefact/text
type ArtiefactObjectPostTextResponse struct {
	ArtiefactObject *ArtiefactObject `json:"artiefact_object,omitempty"`
}

// ArtiefactUserGetUserResponse struct for artiefact_user
// GET: /get-user
type ArtiefactUserGetUserResponse struct {
	ArtiefactUser *ArtiefactUser `json:"artiefact_user,omitempty" schema:"artiefact_user"`
}

// ArtiefactUserSignInRequest struct for artiefact_user
// POST: /sign-in
type ArtiefactUserSignInRequest struct {
	Password string `json:"password,omitempty"`
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
	Password string `json:"password,omitempty"`
	Username string `json:"username"`
}

// ArtiefactUserSignUpResponse struct for artiefact_user
// POST: /sign-up
type ArtiefactUserSignUpResponse struct {
	AccessToken   *AccessToken   `json:"access_token,omitempty"`
	ArtiefactUser *ArtiefactUser `json:"artiefact_user,omitempty"`
}

// ArtiefactUserUsernameAvailabilityRequest struct for artiefact_user
// POST: /username-availability
type ArtiefactUserUsernameAvailabilityRequest struct {
	Username string `json:"username"`
}

// ArtiefactUserUsernameAvailabilityResponse struct for artiefact_user
// POST: /username-availability
type ArtiefactUserUsernameAvailabilityResponse struct {
	IsAvailable bool   `json:"is_available,omitempty"`
	Username    string `json:"username"`
}
