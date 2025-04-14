package http

import (
	"net/http"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/progress"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/upload"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, uploadUsecase *upload.Upload, progressUsecase *progress.Progress) {
	uploadHandler := NewUploadHandler(uploadUsecase)
	progressHandler := NewProgressHandler(progressUsecase)

	r.Post("/upload/chunk", uploadHandler.UploadChunk)
	r.Post("/upload/file", uploadHandler.UploadFile)
	r.Get("/progress", progressHandler.StreamProgress)
	r.Get("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Welcome to api.beambridge.com"}`))
	}))
}
