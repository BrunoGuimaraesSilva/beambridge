package mock

import (
	"context"
	"fmt"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
)

type MockMetadata struct{}

func New() *MockMetadata {
	return &MockMetadata{}
}

func (m *MockMetadata) Extract(ctx context.Context, file *domain.File) (map[string]string, error) {
	return map[string]string{
		"name": file.Name,
		"size": fmt.Sprintf("%d", file.Size),
	}, nil
}
