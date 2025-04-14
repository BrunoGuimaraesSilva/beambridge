package http

import (
	"io"
	"net/http"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/upload"
)

type UploadHandler struct {
	usecase *upload.Upload
}

func NewUploadHandler(usecase *upload.Upload) *UploadHandler {
	return &UploadHandler{usecase: usecase}
}

func (h *UploadHandler) UploadChunk(w http.ResponseWriter, r *http.Request) {
	MAX_CHUNK_SIZE := 32 << 20
	err := r.ParseMultipartForm(int64(MAX_CHUNK_SIZE))
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	fileName := r.FormValue("fileName")
	uploadID := r.FormValue("uploadId")
	index := 0
	fileType := r.FormValue("fileType")

	f, _, err := r.FormFile("chunk")
	if err != nil {
		domain.BadRequestError("Failed to read chunk", "Must be a valid file")
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		domain.BadRequestError("Failed to read chunk data", "Must be a valid file")
		return
	}

	DUMMY_SIZE := int64(len(data) * 2)
	file, err := domain.NewFile(fileName, DUMMY_SIZE)
	if err != nil {
		domain.BadRequestError("Invalid file", "Must be a valid file")
		return
	}

	if fileType == "" {
		err = h.usecase.UploadChunk(r.Context(), uploadID, file, index, data)
	} else {
		livePhoto, err := domain.NewLivePhoto(fileName+".jpg", int64(len(data)), fileName+".mov", int64(len(data)))
		if err != nil {
			domain.BadRequestError("Invalid Live Photo", "Must be a valid file")
			return
		}
		_ = h.usecase.UploadLivePhoto(r.Context(), uploadID, livePhoto, fileType, index, data)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	MAX_FILE_SIZE := 32 << 20
	err := r.ParseMultipartForm(int64(MAX_FILE_SIZE))
	if err != nil {
		domain.BadRequestError("Failed to parse form", "Must be a valid file")
		return
	}

	fileName := r.FormValue("fileName")
	uploadID := r.FormValue("uploadId")
	if fileName == "" || uploadID == "" {
		domain.BadRequestError("Missing fileName or uploadId", "Must be a valid file")
		return
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		domain.BadRequestError("Failed to read file", "Must be a valid file")
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		domain.BadRequestError("Failed to read file data", "Must be a valid file")
		return
	}

	file, err := domain.NewFile(fileName, int64(len(data)))
	if err != nil {
		domain.BadRequestError("Invalid file", "Must be a valid file")
		return
	}

	err = h.usecase.UploadFile(r.Context(), uploadID, file, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
