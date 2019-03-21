package schema

// User struct for user resource
type User struct {
	Birthday     string `json:"birthday,omitempty"`
	Email        string `json:"email"`
	ID           string `json:"id,omitempty"`
	Password     string `json:"password"`
	RegisterDate string `json:"register_date,omitempty"`
	Status       string `json:"status,omitempty"`
	Username     string `json:"username"`
}

// UserGetUserRequest struct for user
// GET: /get-user
type UserGetUserRequest struct {
	Email    string `json:"email" schema:"email"`
	Password string `json:"password" schema:"password"`
	Username string `json:"username" schema:"username"`
}

// UserGetUserResponse struct for user
// GET: /get-user
type UserGetUserResponse struct {
	User *User `json:"user,omitempty" schema:"user"`
}
