package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Connection *sqlx.DB
}

func (db *DB) Open(ctx context.Context, driverName string, conString string, maxConnections int) error {
	if driverName != "mysql" && driverName != "postgres" {
		return errors.New(fmt.Sprintf("invalid driverName: %s %s", driverName, conString))
	}
	conn, err := sqlx.ConnectContext(ctx, driverName, conString)
	if err != nil {
		return err
	}
	conn.SetMaxIdleConns(maxConnections)
	conn.SetMaxOpenConns(maxConnections)
	conn.SetConnMaxLifetime(5 * time.Minute)

	db.Connection = conn

	return db.Connection.Ping()
}

func (db *DB) Close() {
	_ = db.Connection.Close()
}
