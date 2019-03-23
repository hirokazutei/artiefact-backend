package c

import "fmt"

// Database Errors
const (
	ErrorDBFailedToBegin  = "database has failed to begin"
	ErrorDBFailedToCommit = "database has failed to commit"
)

// Common Errors
const (
	ErrorAlreadyExists = "the specified %s already exists"
)

// ErrorAction create an error message with a current actio and object
func ErrorAction(action string, object string) string {
	return fmt.Sprintf("encountered an error while %s %s", action, object)
}

// ErrorActionDetail create an error message with a current action, object, and detail
func ErrorActionDetail(action string, object string, detail string) string {
	return fmt.Sprintf("encountered an error while %s %s: %s", action, object, detail)
}
