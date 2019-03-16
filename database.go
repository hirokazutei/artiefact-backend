package artiefact

import (
	"database/sql"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

// Database Struct
type Database struct {
	DatabaseConnection
}

// QueryMaker Interface
type QueryMaker interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// TransactionMaker Interface
type TransactionMaker interface {
	QueryMaker
	Commit() error
	Rollback() error
}

// DatabaseConnection Interface
type DatabaseConnection interface {
	QueryMaker
	Begin() (*sql.Tx, error)
	Close() error
	Ping() error
}

// NewDatabase create a new DatabaseObject
func NewDatabase(config *AppConfig) (*Database, error) {
	databaseConfigureation := &stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     config.DatabaseHost,
			User:     config.DatabaseUser,
			Password: config.DatabasePass,
			Database: config.DatabaseName,
			Port:     config.DatabasePort,
		},
	}
	stdlib.RegisterDriverConfig(databaseConfigureation)
	database, err := sql.Open("pgx", databaseConfigureation.ConnectionString(""))
	if err != nil {
		return nil, err
	}
	database.SetMaxOpenConns(20)
	database.SetMaxIdleConns(10)
	// db.SetConnMaxLifetime(time.Second * 10)
	return &Database{database}, nil
}
