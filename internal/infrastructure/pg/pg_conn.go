package pg

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGDBInstance struct {
	*pgxpool.Pool
}

func NewPool() (*PGDBInstance, error) {
	connectionString := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), connectionString)

	if err != nil {
		return nil, err
	}

	return &PGDBInstance{
		dbpool,
	}, nil
}