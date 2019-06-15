package artiefact

import (
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
)

// UserApp is app for User resource
type UserApp struct {
	*App
}

// UserHandler user handler
type UserHandler struct {
	handler func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

// ServeHTTPC for user
func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	status, res, err := h.handler(w, r)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(status)
	if err := encoder.Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(c.ErrorUnknown)
		return
	}
	return
}

// SignUpHandler create user and return token
func (app *UserApp) SignUpHandler(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var param schema.UserSignupRequest

	// Validate the JSON coming in with the appropriate JSON-Schema Validator
	res, err := schema.Validate(&param, schema.UserSignupValidator, r)
	fmt.Println(res) // kaz
	if err != nil {
		// json.NewEncoder(w).Encode(res)
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("validating", "request"))
		fmt.Println(err.Error()) // kaz
		return e.GenerateResponse()
	}

	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := c.ErrorDatabaseBeginFailure()
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}
	defer tx.Rollback()

	// Check if Email is taken
	emailExists, err := model.IfEmailExist(tx, param.Email)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("querying", "email"))
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}
	if emailExists {
		e := c.ErrorAlreadyExists("email")
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}

	// TODO: Email Verification
	// Unverified email should restrict features, not prevent users from accessing basic features

	// Check if Username is taken
	usernameExists, err := model.IfUsernameExist(tx, param.Username)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("querying", "username"))
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}
	if usernameExists {
		e := c.ErrorAlreadyExists("username")
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}

	/*
		Season your passwords,
		Not 'cause of security,
		But for better taste.
	*/

	// Pepper Password
	hashedPassword, err := service.PepperAndSaltPassward(param.Password, app.Config.PasswordPepper)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("generating", "password"))
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}

	// Convert Birthday
	birthday, err := time.Parse(c.DateFormat, param.Birthday)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("parsing", "birthday"))
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
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
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("creating", "artiefact_user"))
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}

	// Create Other User Profile Too
	newUsername := model.Username{
		UserID:        newUser.ID,
		UsernameLower: strings.ToLower(param.Username),
		UsernameRaw:   param.Username,
	}
	err = newUsername.Create(tx)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("creating", "username"))
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
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
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("creating", "token"))
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}

	err = tx.Commit()
	if err != nil {
		e := c.ErrorDatabaseCommitFailure()
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(e)
		return e.GenerateResponse()
	}

	// Create Response
	response := schema.UserSignupResponse{
		Token: tokenString,
	}

	// json.NewEncoder(w).Encode(response)
	return http.StatusOK, response, nil
}

// SignInHandler log-in user and create
func (app *UserApp) SignInHandler(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var param schema.UserSignupRequest

	// Validate the JSON coming in with the appropriate JSON-Schema Validator
	res, err := schema.Validate(&param, schema.UserSigninValidator, r)
	fmt.Println(res)

	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("generating", "password"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}

	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := c.ErrorDatabaseBeginFailure()
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	defer tx.Rollback()

	au, err := model.GetArtiefactUserByUsername(tx, param.Username)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("querying", "artiefact_user"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	if au == nil {
		e := c.ErrorObjectNotFound("artiefact_user")
		return e.GenerateResponse()
	}

	match, err := service.AuthenticatePassword(param.Password, au.Password, app.Config.PasswordPepper)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("authenticating", "password"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	if !match {
		e := c.ErrorAuthenticationFailure()
		return e.GenerateResponse()
	}

	// Generate Token
	tokenGeneratedDatetime := time.Now()
	tokenExpiryDatetime := tokenGeneratedDatetime.AddDate(1, 0, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         au.ID,
		"expiry_datetime": tokenExpiryDatetime,
		"obtained_by":     c.TokenObtainedBySignup,
		"tokenType":       c.TokenTypeLogin,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(app.Config.TokenSecret))

	newToken := model.AccessToken{
		Token:             tokenString,
		UserID:            au.ID,
		GeneratedDatetime: tokenGeneratedDatetime,
		ExpiryDatetime:    tokenExpiryDatetime,
		ObtainedBy:        c.TokenObtainedBySignup,
		TokenType:         c.TokenTypeLogin,
	}

	err = newToken.Create(tx)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("creating", "token"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}

	err = tx.Commit()
	if err != nil {
		e := c.ErrorDatabaseCommitFailure()
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}

	// Create Response
	response := schema.UserSignupResponse{
		Token: tokenString,
	}

	// json.NewEncoder(w).Encode(response)
	return http.StatusOK, response, nil
}
