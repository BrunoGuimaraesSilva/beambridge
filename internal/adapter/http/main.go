package http

import (
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/auth"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/progress"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/upload"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, uploadUsecase *upload.Upload, progressUsecase *progress.Progress, authUsecase *auth.Auth, jwtSecret string) {
	uploadHandler := NewUploadHandler(uploadUsecase)
	progressHandler := NewProgressHandler(progressUsecase)
	authHandler := NewAuthHandler(authUsecase)
	r.Use(JSONMiddleware)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Get("/refresh", authHandler.RefreshToken)
	})

	r.Route("/upload", func(r chi.Router) {
		r.Use(AuthMiddleware(jwtSecret))
		r.Post("/chunk", uploadHandler.UploadChunk)
		r.Post("/file", uploadHandler.UploadFile)
	})

	r.Get("/progress", progressHandler.StreamProgress)
}
