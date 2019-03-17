package artiefact

import (
	"database/sql"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

// DB Struct
type DB struct {
	DBer
}

// Queryer Interface
type Queryer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// TransactionMaker Interface
type TransactionMaker interface {
	Queryer
	Commit() error
	Rollback() error
}

// DBer Interface
type DBer interface {
	Queryer
	Begin() (*sql.Tx, error)
	Close() error
	Ping() error
}

// NewDatabase create a new DatabaseObject
func NewDatabase(config *AppConfig) (*DB, error) {
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
	return &DB{database}, nil
}
