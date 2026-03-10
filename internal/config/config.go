package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	S3Endpoint  string
	S3AccessKey string
	S3SecretKey string
	S3Bucket    string

	ImagesDir   string
	DownloadDir string
}

func Load() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system env")
	}

	return &Config{
		S3Endpoint:  os.Getenv("S3_ENDPOINT"),
		S3AccessKey: os.Getenv("S3_ACCESS_KEY"),
		S3SecretKey: os.Getenv("S3_SECRET_KEY"),
		S3Bucket:    os.Getenv("S3_BUCKET"),
		ImagesDir:   os.Getenv("IMAGES_DIR"),
		DownloadDir: os.Getenv("DOWNLOAD_DIR"),
	}
}
