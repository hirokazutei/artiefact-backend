package schema

// Token struct for token resource
type Token struct {
	Status string `json:"status,omitempty"`
	Token  string `json:"token,omitempty"`
}

// User struct for user resource
type User struct {
	Birthday     string `json:"birthday"`
	Email        string `json:"email"`
	ID           string `json:"id,omitempty"`
	Password     string `json:"password"`
	RegisterDate string `json:"register_date,omitempty"`
	Status       string `json:"status,omitempty"`
	Username     string `json:"username"`
}

// UserSigninRequest struct for user
// POST: /signin
type UserSigninRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// UserSigninResponse struct for user
// POST: /signin
type UserSigninResponse struct {
	Token string `json:"token,omitempty"`
}

// UserSignupRequest struct for user
// POST: /signup
type UserSignupRequest struct {
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// UserSignupResponse struct for user
// POST: /signup
type UserSignupResponse struct {
	Token string `json:"token,omitempty"`
}
