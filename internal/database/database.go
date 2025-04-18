package database

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewDatabase(dsn string, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, status.Error(codes.Internal, "Can't open database")
	}
	return db, nil
}
