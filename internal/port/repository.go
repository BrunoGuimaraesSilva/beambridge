package port

import (
	"context"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
)

type FileRepository interface {
	SaveFile(ctx context.Context, file *domain.File) error
	SaveLivePhoto(ctx context.Context, livePhoto *domain.LivePhoto) error
}

type SessionRepository interface {
	UpdateProgress(ctx context.Context, uploadID string, percent int) error
	GetProgress(ctx context.Context, uploadID string) (int, error)
}
