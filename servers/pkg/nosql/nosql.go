package nosql

import "context"

// Operator is Range condition.
type Operator string

// NoSQL is the interface implemented by an object
// that can handle NoSQL.
type NoSQL interface {
	Table(string) Table
}

// Table acts a table object in NoSQL.
type Table interface {
	Get(string, interface{}) Query
	Put(interface{}) Put
	Delete(string, interface{}) Delete
}

// Query represents a query to get items.
type Query interface {
	All(interface{}) error
	AllWithContext(context.Context, interface{}) error
	OneWithContext(context.Context, interface{}) error
	Index(string) Query
	Range(string, Operator, ...interface{}) Query
	Filter(string, ...interface{}) Query
}

// Put represents a query to put item.
type Put interface {
	Run() error
	RunWithContext(context.Context) error
}

// Delete represents a query to delete items.
type Delete interface {
	Range(string, interface{}) Delete
	If(string, ...interface{}) Delete
	Run() error
	RunWithContext(context.Context) error
}
