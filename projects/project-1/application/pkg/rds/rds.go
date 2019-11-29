package rds

import (
	"context"
)

// RDS is the interface implemented by the object that
// acts the MySQL client.
type RDS interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// Result is the interface implemented by the object that
// represents the result of MySQL.
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Rows is the interface implemented by the object that represents
// the result rows of MySQL.
type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
