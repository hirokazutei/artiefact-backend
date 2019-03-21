package schema

//go:generate gom exec prmdg struct --file=../json-schema/schema/schema.json --package=schema --output=./struct.go
//go:generate gom exec prmdg jsval --file=../json-schema/schema/schema.json --package=schema --output=./validator.go
