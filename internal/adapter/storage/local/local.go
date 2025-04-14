package local

import (
	"context"
	"os"
	"path/filepath"

	"fmt"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
)

type LocalStorage struct {
	basePath string
}

func New(basePath string) *LocalStorage {
	os.MkdirAll(basePath, 0755)
	return &LocalStorage{basePath: basePath}
}

func (s *LocalStorage) SaveChunk(ctx context.Context, fileID string, index int, data []byte) error {
	path := filepath.Join(s.basePath, fileID, "chunk"+fmt.Sprint(index))
	os.MkdirAll(filepath.Dir(path), 0755)
	return os.WriteFile(path, data, 0644)
}

func (s *LocalStorage) AssembleFile(ctx context.Context, file *domain.File) (string, error) {
	path := filepath.Join(s.basePath, file.ID.String(), file.Name)
	os.MkdirAll(filepath.Dir(path), 0755)
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	for i := range file.ChunkCount {
		data, ok := file.Chunks[i]
		if !ok {
			return "", os.ErrNotExist
		}
		if _, err := f.Write(data); err != nil {
			return "", err
		}
	}
	return path, nil
}

func (s *LocalStorage) SaveFile(ctx context.Context, fileID string, fileName string, data []byte) (string, error) {
	path := filepath.Join(s.basePath, fileID, fileName)
	os.MkdirAll(filepath.Dir(path), 0755)
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", err
	}
	return path, nil
}
