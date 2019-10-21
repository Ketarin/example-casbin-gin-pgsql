package infra

import (
	"github.com/jackc/pgx"
)

//Connection mockable requester object
type Connection struct {
	Conn *pgx.ConnPool
}

//ConnectionInterface used to available mock for Connection
type ConnectionInterface interface {
	Query(sql string, args ...interface{}) (*pgx.Rows, error)
}

//Query wrapp pqx.Query method
func (c Connection) Query(sql string, args ...interface{}) (*pgx.Rows, error) {
	return c.Conn.Query(sql, args...)
}
