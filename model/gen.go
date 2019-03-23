package model

//go:generate gom exec dgw postgres://artiefact@localhost/artiefact?sslmode=disable --schema=artiefact --package=model --output=table.go --exclude=alembic_version
