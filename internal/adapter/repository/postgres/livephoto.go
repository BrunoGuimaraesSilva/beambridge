package postgres

import (
	"context"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
)

func (r *PostgresRepo) SaveLivePhoto(ctx context.Context, livePhoto *domain.LivePhoto) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := r.SaveFile(ctx, livePhoto.ImageFile); err != nil {
		return err
	}
	if err := r.SaveFile(ctx, livePhoto.VideoFile); err != nil {
		return err
	}
	query := `INSERT INTO live_photos (id, image_file_id, video_file_id, metadata) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(ctx, query, livePhoto.ID, livePhoto.ImageFile.ID, livePhoto.VideoFile.ID, livePhoto.Metadata)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
