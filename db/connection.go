package db

import (
	"context"
	"os"
	"time"

	"fmt"

	"github.com/jmoiron/sqlx"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

var (
	db *sqlx.DB
	// Mock ...
	Mock sqlxmock.Sqlmock
	// InitMocked ...
	InitMocked bool
)

// GetConnection ...
func GetConnection() (_ *sqlx.DB, err error) {
	if db == nil {
		if InitMocked {
			db, Mock, err = sqlxmock.Newx()
			if err != nil {
				return nil, err
			}
			return
		}

		if db, err = sqlx.Connect(
			"postgres",
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"),
				os.Getenv("DB_NAME"),
			),
		); err != nil {
			return nil, err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			return nil, fmt.Errorf("%v: %v", "Failed trying to ping database", err)
		}
	}

	return db, nil
}
