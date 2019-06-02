package model

//go:generate gom exec dgw postgres://artiefact@localhost/artiefact?sslmode=disable --package=model --output=table.go --exclude=alembic_version

// Remember that you can specify schema.
