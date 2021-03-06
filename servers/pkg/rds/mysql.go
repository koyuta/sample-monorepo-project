package rds

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	// mysql driver is used only this file
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// MySQL acts as the MySQL client.
type MySQL struct {
	conn *sqlx.DB
}

// Connect returns a new MySQL.
func Connect(user, passwd, host string, port uint, dbname string) (*MySQL, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, passwd, host, port, dbname)
	conn, err := sqlx.Connect("mysql", uri)
	if err != nil {
		return nil, err
	}
	return &MySQL{conn: conn}, err
}

// ConnectWithContext returns a new MySQL.
func ConnectWithContext(ctx context.Context, user, passwd, host string, port uint, dbname string) (*MySQL, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, passwd, host, port, dbname)
	conn, err := sqlx.ConnectContext(ctx, "mysql", uri)
	if err != nil {
		return nil, err
	}
	return &MySQL{conn: conn}, err
}

// ExecContext execute a query via Conn.ExecContext function. It expects
// to use CREATE, UPDATE, DELETE commands.
func (m *MySQL) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error) {
	res, err := m.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return new(MySQLResult), err
	}
	return &MySQLResult{Result: res}, nil
}

// GetContext executes a query and binds a result raw data to dest that passed.
func (m *MySQL) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := m.conn.GetContext(ctx, dest, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}

// SelectContext executes a query and binds a result raw data to dest that passwd.
func (m *MySQL) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.conn.SelectContext(ctx, dest, query, args...)
}

// MySQLResult represents the result of a query.
type MySQLResult struct {
	Result sql.Result
}

// LastInsertId returns a integer generated by the database
// in response to a command.
func (r *MySQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected returns the number of rows affected.
func (r *MySQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// MySQLRows represents the result of a query.
type MySQLRows struct {
	Rows *sql.Rows
}

// Scan assigns a value from database.
func (r *MySQLRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next prepares the next result row for reading with the Scan.
func (r *MySQLRows) Next() bool {
	return r.Rows.Next()
}

// Close closes the Rows, preventing further enumeration.
func (r *MySQLRows) Close() error {
	return r.Rows.Close()
}
