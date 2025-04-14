package postgres

import (
	"context"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (r *PostgresRepo) SaveUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, email, password_hash) VALUES ($1, $2, $3)`
	_, err := r.conn.Exec(ctx, query, user.ID, user.Email, user.PasswordHash)
	return err
}

func (r *PostgresRepo) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	user := &domain.User{}
	err := r.conn.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
