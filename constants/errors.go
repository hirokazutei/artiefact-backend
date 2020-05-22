package c

import (
	"errors"
	"fmt"
	"net/http"
)

// Error struct for error resource
type Error struct {
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status"`
	Type   string `json:"type"`
}

// Custom Status Codes
const (
	StatusUnknownError = 520
)

// Server Error
const (
	ErrorTypeInternalServer = "server has encountered an error"
)

// Database Errors
const (
	ErrorTypeDBFailedToBegin  = "database has failed to begin"
	ErrorTypeDBFailedToCommit = "database has failed to commit"
	ErrorTypeObjectNotFound   = "the desired object could not be found"
)

// Common Errors
const (
	ErrorTypeObjectAlreadyExists = "the object already exists"
	ErrorTypeInvalidHeader       = "the request header was invalid"
	ErrorTypeInvalidRequest      = "the request was invalid"
	ErrorTypeUnknown             = "the server encountered an unknown error"
)

// Authentication Errors
const (
	ErrorTypeAuthenticationFailure = "could not authenticate with the given information"
	ErrorTypeInvalidToken          = "the token is invalid"
)

// ErrorFunctionFailure creates an error message with the current function
func ErrorFunctionFailure(function string) string {
	return fmt.Sprintf("%s has failed", function)
}

// ErrorAction create an error message with a current action and object
func ErrorAction(action string, object string) string {
	return fmt.Sprintf("encountered an error while %s %s", action, object)
}

// ErrorActionDetail create an error message with a current action, object, and detail
func ErrorActionDetail(action string, object string, detail string) string {
	return fmt.Sprintf("encountered an error while %s %s: %s", action, object, detail)
}

// AddDetail to the Error
func (err Error) AddDetail(detail string) *Error {
	err.Detail = detail
	return &err
}

// GetStatus obtains error status
func (err Error) GetStatus() int {
	return err.Status
}

// GetType obtains error status
func (err Error) GetType() string {
	return err.Type
}

// GetDetail obtains error status
func (err Error) GetDetail() string {
	return err.Detail
}

// GetError returns the error detail
func (err Error) GetError() error {
	errorDetail := err.Detail
	return errors.New(errorDetail)
}

// GenerateResponse returns the necessary parameters for handler functions
func (err Error) GenerateResponse() (int, interface{}, error) {
	return err.GetStatus(), err, err.GetError()
}

// ErrorInvalidHeader returns Invalid Header Error
func ErrorInvalidHeader() *Error {
	return &Error{
		Status: http.StatusBadRequest,
		Type:   ErrorTypeInvalidHeader,
	}
}

// ErrorUnknown returns Unknown Error
func ErrorUnknown() *Error {
	return &Error{
		Status: StatusUnknownError,
		Type:   ErrorTypeUnknown,
	}
}

// ErrorInternalServer returns Internal Server Error
func ErrorInternalServer() *Error {
	return &Error{
		Status: http.StatusInternalServerError,
		Type:   ErrorTypeInternalServer,
	}
}

// ErrorInvalidToken returns Invalid Token Error
func ErrorInvalidToken() *Error {
	return &Error{
		Status: http.StatusBadRequest,
		Type:   ErrorTypeInvalidToken,
	}
}

// ErrorAuthenticationFailure returns Authentication Failure Error
func ErrorAuthenticationFailure() *Error {
	return &Error{
		Status: http.StatusBadRequest,
		Type:   ErrorTypeAuthenticationFailure,
	}
}

// ErrorObjectNotFound returns Object Not Found Error
func ErrorObjectNotFound(object string) *Error {
	return &Error{
		Status: http.StatusNotFound,
		Type:   ErrorTypeObjectNotFound,
		Detail: ErrorMessageObjectNotFound(object),
	}
}

// ErrorDatabaseBeginFailure returns Database Begin Failure Error
func ErrorDatabaseBeginFailure() *Error {
	return &Error{
		Status: http.StatusInternalServerError,
		Type:   ErrorTypeDBFailedToBegin,
	}
}

// ErrorDatabaseCommitFailure returns Database Commit Failure Error
func ErrorDatabaseCommitFailure() *Error {
	return &Error{
		Status: http.StatusInternalServerError,
		Type:   ErrorTypeDBFailedToCommit,
	}
}

// ErrorAlreadyExists returns Already Exists Error
func ErrorAlreadyExists(object string) *Error {
	return &Error{
		Status: http.StatusConflict,
		Type:   ErrorTypeObjectAlreadyExists,
		Detail: ErrorMessageObjectAlreadyExists(object),
	}
}

// ErrorBadRequest returns Bad Request Error
func ErrorBadRequest() *Error {
	return &Error{
		Status: http.StatusBadRequest,
		Type:   ErrorTypeInvalidRequest,
	}
}

// Error Messages

// ErrorMessageObjectAlreadyExists returns Object Already Exists Error Message
func ErrorMessageObjectAlreadyExists(object string) string {
	return fmt.Sprintf("the %s already exists", object)
}

// ErrorMessageObjectNotFound returns Object Not Found Error Message
func ErrorMessageObjectNotFound(object string) string {
	return fmt.Sprintf("the requested %s could not be found", object)
}
