package main

import (
	"log"
	"os"

	"github.com/Ketarin/pg-casbin/app/handler"
	"github.com/Ketarin/pg-casbin/app/infra"
	"github.com/Ketarin/pg-casbin/app/process"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

func main() {

	r := gin.Default()

	conn, err := initDatabase()

	if err != nil {
		log.Fatalf("Database error : %s", err)
	}

	r.GET("/people", handler.Auth(conn), process.SomeProcess())
	r.POST("/people", handler.Auth(conn), process.SomeProcess())
	r.PUT("/people/:id", handler.Auth(conn), process.SomeProcess())
	r.DELETE("/people/:id", handler.Auth(conn), process.SomeProcess())

	r.Run()

}

func initDatabase() (infra.Connection, error) {
	connPool, err := pgx.NewConnPool(
		pgx.ConnPoolConfig{
			ConnConfig: pgx.ConnConfig{
				Host:     os.Getenv("PG_HOST"),
				Database: os.Getenv("PG_DB"),
				User:     os.Getenv("PG_USER"),
				Password: os.Getenv("PG_PASSWORD"),
			},
		},
	)
	if err != nil {
		return infra.Connection{}, err
	}

	return infra.Connection{Conn: connPool}, nil
}
