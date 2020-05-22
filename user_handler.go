package artiefact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	var param schema.ArtiefactUserSignUpRequest

	// Validate the JSON coming in with the appropriate JSON-Schema Validator
	res, err := schema.Validate(&param, schema.ArtiefactUserSignUpValidator, r)
	fmt.Println(res)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("validating", "request"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	fmt.Println(param)

	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := c.ErrorDatabaseBeginFailure()
		return e.GenerateResponse()
	}
	defer tx.Rollback()

	// Check if Email is taken
	emailExists, err := model.IfEmailExist(tx, param.Email)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("querying", "email"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	if emailExists {
		e := c.ErrorAlreadyExists("email")
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
		return e.GenerateResponse()
	}
	if usernameExists {
		e := c.ErrorAlreadyExists("username")
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
		return e.GenerateResponse()
	}

	// Convert Birthday
	birthday, err := time.Parse(c.DateFormat, param.Birthday)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("parsing", "birthday"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}

	// Create the user
	newUser := model.ArtiefactUser{
		Password:         hashedPassword,
		Birthday:         birthday,
		RegisterDatetime: time.Now(),
		Status:           c.UserUnverified,
	}

	err = newUser.Create(tx)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("creating", "artiefact_user"))
		fmt.Println(err.Error())
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
		return e.GenerateResponse()
	}

	// Generate Token
	tokenGeneratedDatetime := time.Now()
	tokenExpiryDatetime := tokenGeneratedDatetime.AddDate(1, 0, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         newUser.ID,
		"expiry_datetime": tokenExpiryDatetime,
		"obtained_by":     c.TokenObtainedBySignup,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(app.Config.TokenSecret))

	newToken := model.AccessToken{
		Token:             tokenString,
		UserID:            newUser.ID,
		GeneratedDatetime: tokenGeneratedDatetime,
		ExpiryDatetime:    tokenExpiryDatetime,
		ObtainedBy:        c.TokenObtainedBySignup,
		Active:            true,
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
	resToken := schema.AccessToken{
		Token:             newToken.Token,
		UserID:            newToken.UserID,
		GeneratedDatetime: newToken.GeneratedDatetime,
		ExpiryDatetime:    newToken.ExpiryDatetime,
		ObtainedBy:        newToken.ObtainedBy,
		Active:            newToken.Active,
	}

	resUser := schema.ArtiefactUser{
		ID:               newUser.ID,
		Birthday:         newUser.Birthday.Format(param.Birthday),
		RegisterDatetime: newUser.RegisterDatetime,
		Status:           newUser.Status,
		Username:         newUsername.UsernameRaw,
	}

	response := schema.ArtiefactUserSignUpResponse{
		AccessToken:   &resToken,
		ArtiefactUser: &resUser,
	}
	return http.StatusOK, response, nil
}

// SignInHandler log-in user and create
func (app *UserApp) SignInHandler(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var param schema.ArtiefactUserSignInRequest

	// Validate the JSON coming in with the appropriate JSON-Schema Validator
	res, err := schema.Validate(&param, schema.ArtiefactUserSignInValidator, r)
	fmt.Println(res)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("validating", "request"))
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
		Active:            true,
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
	resToken := schema.AccessToken{
		Token:             newToken.Token,
		UserID:            newToken.UserID,
		GeneratedDatetime: newToken.GeneratedDatetime,
		ExpiryDatetime:    newToken.ExpiryDatetime,
		ObtainedBy:        newToken.ObtainedBy,
		Active:            newToken.Active,
	}

	resUser := schema.ArtiefactUser{
		ID:               au.ID,
		Birthday:         au.Birthday.Format(c.DateFormat),
		RegisterDatetime: au.RegisterDatetime,
		Status:           au.Status,
		Username:         param.Username,
	}

	response := schema.ArtiefactUserSignUpResponse{
		AccessToken:   &resToken,
		ArtiefactUser: &resUser,
	}

	return http.StatusOK, response, nil
}

// GetUserHandler obtains user associated with the token and returns the user.
func (app *UserApp) GetUserHandler(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := r.Context()
	u, ok := ctx.Value(contextKeyAuth).(*AuthResponse)
	if !ok {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("obtaining", "context value"))
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

	username, found, err := model.GetUsernameByUserID(tx, u.User.ID)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("querying", "username"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	if !found {
		e := c.ErrorObjectNotFound("username")
		return e.GenerateResponse()
	}

	resUser := schema.ArtiefactUser{
		ID:               u.User.ID,
		Birthday:         u.User.Birthday.Format(c.DateFormat),
		RegisterDatetime: u.User.RegisterDatetime,
		Status:           u.User.Status,
		Username:         username.UsernameRaw,
	}

	// TODO: Should Return Other User Info as well

	var response = schema.ArtiefactUserGetUserResponse{
		ArtiefactUser: &resUser,
	}

	return http.StatusOK, response, nil
}

// UsernameAvailabilityHandler check to see if the username is available
func (app *UserApp) UsernameAvailabilityHandler(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		e := c.ErrorBadRequest()
		e.AddDetail(c.ErrorAction("parsing", "request"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	username := m["username"][0]

	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := c.ErrorDatabaseBeginFailure()
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	defer tx.Rollback()

	_, found, err := model.GetUsernameByUsername(tx, username)
	if err != nil {
		e := c.ErrorInternalServer()
		e.AddDetail(c.ErrorAction("querying", "username"))
		fmt.Println(err.Error())
		return e.GenerateResponse()
	}
	var response = schema.ArtiefactUserUsernameAvailabilityResponse{
		Username:    username,
		IsAvailable: false,
	}

	if !found {
		response.IsAvailable = true
	}
	// Intentionally taking extra time
	time.Sleep(2 * time.Second)
	return http.StatusOK, response, nil
}
