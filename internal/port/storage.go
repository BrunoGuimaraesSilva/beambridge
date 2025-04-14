package port

import (
	"context"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
)

type Storage interface {
	SaveChunk(ctx context.Context, fileID string, index int, data []byte) error
	AssembleFile(ctx context.Context, file *domain.File) (string, error)
	SaveFile(ctx context.Context, fileID string, fileName string, data []byte) (string, error)
}
