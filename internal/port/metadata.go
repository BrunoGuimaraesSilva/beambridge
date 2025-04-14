package port

import (
	"context"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
)

type MetadataExtractor interface {
	Extract(ctx context.Context, file *domain.File) (map[string]string, error)
}
