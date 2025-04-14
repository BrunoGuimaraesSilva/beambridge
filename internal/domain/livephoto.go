package domain

import "github.com/google/uuid"

type LivePhoto struct {
	ID        uuid.UUID
	ImageFile *File
	VideoFile *File
	Metadata  map[string]string
}

func NewLivePhoto(imageName string, imageSize int64, videoName string, videoSize int64) (*LivePhoto, error) {
	imageFile, err := NewFile(imageName, imageSize)
	if err != nil {
		return nil, err
	}
	videoFile, err := NewFile(videoName, videoSize)
	if err != nil {
		return nil, err
	}
	return &LivePhoto{
		ID:        uuid.New(),
		ImageFile: imageFile,
		VideoFile: videoFile,
		Metadata:  make(map[string]string),
	}, nil
}
