package app

import (
	"context"
	"log"
	"s3/internal/config"
	"s3/internal/service"
	"s3/internal/storage/s3"
)

func Run() {
	ctx := context.Background()
	cfg := config.Load()
	s3Client, err := s3.NewClient(cfg)
	if err != nil {
		log.Fatalf("application stopped with error: %v", err)
	}
	imageService := service.NewImageService(s3Client)
	if err := imageService.UploadFromFolder(ctx, cfg.ImagesDir); err != nil {
		log.Fatalf("cant upload pic: %v", err)
	}
	if err := imageService.DownloadToFolder(ctx, cfg.DownloadDir); err != nil {
		log.Fatalf("cant download pic: %v", err)
	}
}
