package artiefact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	c "github.com/hirokazu/artiefact-backend/constants"
	"github.com/hirokazu/artiefact-backend/model"
	"github.com/hirokazu/artiefact-backend/schema"
	"github.com/hirokazu/artiefact-backend/service"
	"golang.org/x/crypto/bcrypt"
)

// UserApp is app for User resource
type UserApp struct {
	*App
}

// Error struct for error resource
type Error struct {
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status"`
	Type   string `json:"type"`
}

// SignUpHandler create user and return token
func (app *UserApp) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var param schema.UserSignupRequest

	// Validate the JSON coming in with the appropriate JSON-Schema Validator
	res, err := schema.Validate(&param, schema.UserSignupValidator, r)
	fmt.Println(res)
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}

	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToBegin,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	defer tx.Rollback()

	// Check if Email is taken
	emailExists, err := model.IfEmailExist(tx, param.Email)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("querying", "email"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	if emailExists {
		e := &Error{
			Status: http.StatusBadRequest,
			Type:   fmt.Sprintf(c.ErrorAlreadyExists, "email"),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// TODO: Email Verification
	// Unverified email should restrict features, not prevent users from accessing basic features

	// Check if Username is taken
	usernameExists, err := model.IfUsernameExist(tx, param.Username)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("querying", "username"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	if usernameExists {
		e := &Error{
			Status: http.StatusBadRequest,
			Type:   fmt.Sprintf(c.ErrorAlreadyExists, "username"),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	/*
		Season your passwords,
		Not 'cause of security,
		But for better taste.
	*/

	// Pepper Password
	hashedPassword, err := service.PepperAndSaltPassward(param.Password, app.Config.PasswordPepper)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("generating", "password"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Convert Birthday
	birthday, err := time.Parse(c.DateFormat, param.Birthday)
	if err != nil {
		e := &Error{
			Status: http.StatusUnprocessableEntity,
			Type:   c.ErrorActionDetail("parsing", "birthday", param.Birthday),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create the user
	newUser := model.ArtiefactUser{
		Password:         hashedPassword,
		Email:            param.Email,
		Birthday:         birthday,
		RegisterDatetime: time.Now(),
		Status:           c.UserUnverified,
	}

	err = newUser.Create(tx)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("creating", "artiefact_user"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create Other User Profile Too
	newUsername := model.Username{
		UserID:        newUser.ID,
		UsernameLower: strings.ToLower(param.Username),
		UsernameRaw:   param.Username,
	}
	err = newUsername.Create(tx)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("creating", "username"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Generate Token
	tokenGeneratedDatetime := time.Now()
	tokenExpiryDatetime := tokenGeneratedDatetime.AddDate(1, 0, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         newUser.ID,
		"expiry_datetime": tokenExpiryDatetime,
		"obtained_by":     c.TokenObtainedBySignup,
		"tokenType":       c.TokenTypeLogin,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(app.Config.TokenSecret))

	newToken := model.AccessToken{
		Token:             tokenString,
		UserID:            newUser.ID,
		GeneratedDatetime: tokenGeneratedDatetime,
		ExpiryDatetime:    tokenExpiryDatetime,
		ObtainedBy:        c.TokenObtainedBySignup,
		TokenType:         c.TokenTypeLogin,
	}

	err = newToken.Create(tx)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("creating", "token"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	err = tx.Commit()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToCommit,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create Response
	response := schema.UserSignupResponse{
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(response)
}

// SignInHandler log-in user and create
func (app *UserApp) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var param schema.UserSignupRequest

	// Validate the JSON coming in with the appropriate JSON-Schema Validator
	res, err := schema.Validate(&param, schema.UserSigninValidator, r)
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}

	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToBegin,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	defer tx.Rollback()

	au, err := model.GetArtiefactUserByUsername(tx, param.Username)
	if err != nil {
		e := &Error{
			Status: http.StatusBadRequest,
			Type:   c.ErrorUserNotFound,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	match, err := service.AuthenticatePassword(param.Password, au.Password, app.Config.PasswordPepper)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type: c.ErrorFunctionFailure("AuthenticatePassword")
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	if !match {
		e := &Error{
			Status: http.StatusBadRequest,
			Type: c.ErrorWrongPassword,
		}
		json.NewEncoder(w).Encode(e)
		return	
	}


	// Pepper Password
	var pepperedPassword bytes.Buffer
	pepperedPassword.WriteString(param.Password)
	pepperedPassword.WriteString(app.Config.PasswordPepper)

	// Generate Salt & Hash
	hashByte, err := bcrypt.GenerateFromPassword([]byte(pepperedPassword.String()), bcrypt.DefaultCost)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("generating", "password"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// See if there are valid tokens

	// Generate Token
	tokenGeneratedDatetime := time.Now()
	tokenExpiryDatetime := tokenGeneratedDatetime.AddDate(1, 0, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         newUser.ID,
		"expiry_datetime": tokenExpiryDatetime,
		"obtained_by":     c.TokenObtainedBySignup,
		"tokenType":       c.TokenTypeLogin,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(app.Config.TokenSecret))

	newToken := model.AccessToken{
		Token:             tokenString,
		UserID:            newUser.ID,
		GeneratedDatetime: tokenGeneratedDatetime,
		ExpiryDatetime:    tokenExpiryDatetime,
		ObtainedBy:        c.TokenObtainedBySignup,
		TokenType:         c.TokenTypeLogin,
	}

	err = newToken.Create(tx)
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorAction("creating", "token"),
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	err = tx.Commit()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToCommit,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create Response
	response := schema.UserSignupResponse{
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(response)
}
