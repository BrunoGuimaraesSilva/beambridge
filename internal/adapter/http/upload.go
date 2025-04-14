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
	err := r.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	fileName := r.FormValue("fileName")
	uploadID := r.FormValue("uploadId")
	index := 0                          // Simplified: parse from form if needed
	fileType := r.FormValue("fileType") // "image" or "video" for Live Photos

	f, _, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, "Failed to read chunk", http.StatusBadRequest)
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, "Failed to read chunk data", http.StatusBadRequest)
		return
	}

	// Mock file creation (in reality, fetch from session or DB)
	file, err := domain.NewFile(fileName, int64(len(data)*2)) // Dummy size
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	if fileType == "" {
		// Regular file upload
		err = h.usecase.UploadChunk(r.Context(), uploadID, file, index, data)
	} else {
		// Live Photo upload
		livePhoto, err := domain.NewLivePhoto(fileName+".jpg", int64(len(data)), fileName+".mov", int64(len(data)))
		if err != nil {
			http.Error(w, "Invalid Live Photo", http.StatusBadRequest)
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
	err := r.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	fileName := r.FormValue("fileName")
	uploadID := r.FormValue("uploadId")
	if fileName == "" || uploadID == "" {
		http.Error(w, "Missing fileName or uploadId", http.StatusBadRequest)
		return
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusBadRequest)
		return
	}

	file, err := domain.NewFile(fileName, int64(len(data)))
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	err = h.usecase.UploadFile(r.Context(), uploadID, file, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
