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
		return nil, ValidationError("invalid file name or size", "file name and size must be provided")
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
		return ValidationError("invalid chunk index", fmt.Sprintf("chunk index must be between 0 and %d", f.ChunkCount-1))
	}
	if len(data) > ChunkSize {
		return ValidationError("chunk size exceeded", fmt.Sprintf("chunk size must not exceed %d bytes", ChunkSize))
	}
	f.Chunks[index] = data
	return nil
}

func (f *File) IsComplete() bool {
	return len(f.Chunks) == f.ChunkCount
}
