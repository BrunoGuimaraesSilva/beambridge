package main

import (
	"log"
	"net/http"

	"github.com/BrunoGuimaraesSilva/beambridge/config"
	httpAdapter "github.com/BrunoGuimaraesSilva/beambridge/internal/adapter/http"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/adapter/metadata/mock"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/adapter/repository/postgres"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/adapter/repository/redis"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/adapter/storage/local"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/progress"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/upload"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	pgRepo, err := postgres.New(cfg.PostgresURL)
	if err != nil {
		log.Fatal("Failed to init PostgreSQL:", err)
	}
	redisRepo := redis.New(cfg.RedisAddr)
	storage := local.New("./uploads")
	metadata := mock.New()

	uploadUsecase := upload.New(pgRepo, redisRepo, storage, metadata)
	progressUsecase := progress.New(redisRepo)

	router := chi.NewRouter()
	httpAdapter.SetupRoutes(router, uploadUsecase, progressUsecase)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed:", err)
	}
}
