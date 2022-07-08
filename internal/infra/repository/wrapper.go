package repository

import (
	"github.com/gocql/gocql"
	"github.com/micheltank/url-shortener-api/internal/infra/cassandra"
	"github.com/sirupsen/logrus"
)

// Wrapper is a wrapper around gocql with a set of convenience functions
// for common operations
type Wrapper struct {
	Cassandra *cassandra.Client
}

// Query wraps makes a query gocql can understand, adding a log message for auditing
func (r *Wrapper) Query(queryString string, values ...interface{}) gocql.Query {
	query := r.Cassandra.Session.Query(queryString, values...)
	logrus.Debug("Executing query: ", query.String())

	return *query
}
