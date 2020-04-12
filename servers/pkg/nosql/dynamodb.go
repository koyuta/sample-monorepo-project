package nosql

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/guregu/dynamo"
)

// DynamoOperator is an alias for Operator.
type DynamoOperator = Operator

// Operation types.
const (
	Equal          DynamoOperator = "EQ"
	NotEqual                      = "NE"
	Less                          = "LT"
	LessOrEqual                   = "LE"
	Greater                       = "GT"
	GreaterOrEqual                = "GE"
	BeginsWith                    = "BEGINS_WITH"
	Between                       = "BETWEEN"
)

// DynamoDB acts as a dynamodb client.
type DynamoDB struct {
	dynamodb *dynamo.DB
}

// NewDynamoDB returns a new DynamoDB.
func NewDynamoDB(p client.ConfigProvider, cfgs ...*aws.Config) *DynamoDB {
	return &DynamoDB{
		dynamodb: dynamo.New(p, cfgs...),
	}
}

// Table returns a table.
func (d *DynamoDB) Table(name string) Table {
	return &DynamoTable{table: d.dynamodb.Table(name)}
}

// DynamoTable represents a table of DynamoDB.
type DynamoTable struct {
	table dynamo.Table
}

// Get creates a query to get items.
func (d *DynamoTable) Get(name string, value interface{}) Query {
	return &DynamoQuery{query: d.table.Get(name, value)}
}

// Put creates a query to put a item.
func (d *DynamoTable) Put(item interface{}) Put {
	return &DynamoPut{put: d.table.Put(item)}
}

// Delete createsa query to delete items.
func (d *DynamoTable) Delete(name string, value interface{}) Delete {
	return &DynamoDelete{del: d.table.Delete(name, value)}
}

// DynamoQuery is a request to get items.
type DynamoQuery struct {
	query *dynamo.Query
}

// All requests this query.
func (d *DynamoQuery) All(out interface{}) error {
	return d.query.All(out)
}

// AllWithContext requests this query with context.
func (d *DynamoQuery) AllWithContext(ctx context.Context, out interface{}) error {
	return d.query.AllWithContext(ctx, out)
}

// OneWithContext requests this query with context.
func (d *DynamoQuery) OneWithContext(ctx context.Context, out interface{}) error {
	return d.query.OneWithContext(ctx, out)
}

// Index sets a secondary index.
func (d *DynamoQuery) Index(name string) Query {
	d.query = d.query.Index(name)
	return d
}

// Filter sets a filterexpression to Query.
func (d *DynamoQuery) Filter(expr string, args ...interface{}) Query {
	d.query = d.query.Filter(expr, args...)
	return d
}

// Range sets a search condition for the RangeKey to query.
func (d *DynamoQuery) Range(name string, op Operator, values ...interface{}) Query {
	d.query = d.query.Range(name, dynamo.Operator(op), values...)
	return d
}

// DynamoPut is a put query.
type DynamoPut struct {
	put *dynamo.Put
}

// Run executes a put query.
func (d *DynamoPut) Run() error {
	return d.put.Run()
}

// RunWithContext executes a put query.
func (d *DynamoPut) RunWithContext(ctx context.Context) error {
	return d.put.RunWithContext(ctx)
}

// DynamoDelete is a delete query.
type DynamoDelete struct {
	del *dynamo.Delete
}

// Range sets a search condtion for the RangeKey to query.
func (d *DynamoDelete) Range(expr string, value interface{}) Delete {
	d.del = d.del.Range(expr, value)
	return d
}

// If sets a condition to this delete query.
func (d *DynamoDelete) If(expr string, args ...interface{}) Delete {
	d.del = d.del.If(expr, args...)
	return d
}

// Run executes a delete query.
func (d *DynamoDelete) Run() error {
	return d.del.Run()
}

// RunWithContext executes a delete query.
func (d *DynamoDelete) RunWithContext(ctx context.Context) error {
	return d.del.RunWithContext(ctx)
}
