package c

// Database Errors
const (
	ErrorDBFailedToBegin  = "database has failed to begin"
	ErrorDBFailedToCommit = "database has failed to commit"
)

// Common Errors
const (
	ErrorAlreadyExists = "the specified %s already exists"
)

// Actions Errors
const (
	ErrorCreating     = "encountered an error while creating %s"
	ErrorCreatingFrom = "encountered an error while creating %s from %s"
	ErrorGenerating   = "encountered an error while generating %s"
	ErrorOpening      = "encountered an error while opening %s"
	ErrorParsingAs    = "encountered an error while parsing %s as %s"
	ErrorReading      = "encountered an error while reading %s"
	ErrorQuerying     = "encountered an error while querying %s"
)
