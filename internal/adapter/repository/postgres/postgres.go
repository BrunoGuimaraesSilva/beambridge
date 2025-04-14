package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresRepo struct {
	conn *pgx.Conn
}

func New(url string) (*PostgresRepo, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepo{conn: conn}, nil
}

func (r *PostgresRepo) Close() {
	r.conn.Close(context.Background())
}
