package postgres

import (
	"context"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
)

func (r *PostgresRepo) SaveFile(ctx context.Context, file *domain.File) error {
	query := `INSERT INTO files (id, name, size, metadata) VALUES ($1, $2, $3, $4)`
	_, err := r.conn.Exec(ctx, query, file.ID, file.Name, file.Size, file.Metadata)
	return err
}
