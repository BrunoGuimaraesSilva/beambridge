package domain

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	ChunkSize = 5 * 1024 * 1024 // 5MB
)

type File struct {
	ID         uuid.UUID
	Name       string
	Size       int64
	ChunkCount int
	Chunks     map[int][]byte
	Metadata   map[string]string
}

func NewFile(name string, size int64) (*File, error) {
	if name == "" || size <= 0 {
		return nil, fmt.Errorf("invalid file name or size")
	}
	chunkCount := int((size + ChunkSize - 1) / ChunkSize)
	return &File{
		ID:         uuid.New(),
		Name:       name,
		Size:       size,
		ChunkCount: chunkCount,
		Chunks:     make(map[int][]byte),
		Metadata:   make(map[string]string),
	}, nil
}

func (f *File) AddChunk(index int, data []byte) error {
	if index < 0 || index >= f.ChunkCount {
		return fmt.Errorf("invalid chunk index")
	}
	if len(data) > ChunkSize {
		return fmt.Errorf("chunk size exceeds limit")
	}
	f.Chunks[index] = data
	return nil
}

func (f *File) IsComplete() bool {
	return len(f.Chunks) == f.ChunkCount
}
