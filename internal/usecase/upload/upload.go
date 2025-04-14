package upload

import (
	"context"

	"maps"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/port"
)

type Upload struct {
	fileRepo    port.FileRepository
	sessionRepo port.SessionRepository
	storage     port.Storage
	metadata    port.MetadataExtractor
}

func New(fileRepo port.FileRepository, sessionRepo port.SessionRepository, storage port.Storage, metadata port.MetadataExtractor) *Upload {
	return &Upload{
		fileRepo:    fileRepo,
		sessionRepo: sessionRepo,
		storage:     storage,
		metadata:    metadata,
	}
}

func (u *Upload) UploadChunk(ctx context.Context, uploadID string, file *domain.File, index int, data []byte) error {
	if err := file.AddChunk(index, data); err != nil {
		return err
	}
	if err := u.storage.SaveChunk(ctx, file.ID.String(), index, data); err != nil {
		return err
	}
	progress := (len(file.Chunks) * 100) / file.ChunkCount
	if err := u.sessionRepo.UpdateProgress(ctx, uploadID, progress); err != nil {
		return err
	}
	if file.IsComplete() {
		path, err := u.storage.AssembleFile(ctx, file)
		if err != nil {
			return err
		}
		file.Metadata["path"] = path
		meta, err := u.metadata.Extract(ctx, file)
		if err != nil {
			return err
		}
		maps.Copy(file.Metadata, meta)
		if err := u.fileRepo.SaveFile(ctx, file); err != nil {
			return err
		}
	}
	return nil
}

func (u *Upload) UploadFile(ctx context.Context, uploadID string, file *domain.File, data []byte) error {
	path, err := u.storage.SaveFile(ctx, file.ID.String(), file.Name, data)
	if err != nil {
		return err
	}
	file.Metadata["path"] = path

	if err := u.sessionRepo.UpdateProgress(ctx, uploadID, 100); err != nil {
		return err
	}

	meta, err := u.metadata.Extract(ctx, file)
	if err != nil {
		return err
	}
	for k, v := range meta {
		file.Metadata[k] = v
	}
	if err := u.fileRepo.SaveFile(ctx, file); err != nil {
		return err
	}

	return nil
}

func (u *Upload) UploadLivePhoto(ctx context.Context, uploadID string, livePhoto *domain.LivePhoto, fileType string, index int, data []byte) error {
	targetFile := livePhoto.ImageFile
	if fileType == "video" {
		targetFile = livePhoto.VideoFile
	}
	if err := u.UploadChunk(ctx, uploadID, targetFile, index, data); err != nil {
		return err
	}
	if livePhoto.ImageFile.IsComplete() && livePhoto.VideoFile.IsComplete() {
		return u.fileRepo.SaveLivePhoto(ctx, livePhoto)
	}
	return nil
}
