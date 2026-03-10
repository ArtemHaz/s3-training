package service

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"io"
	"os"
)
type ImageService struct {
	storage ImageStorage
}

func NewImageService(storage ImageStorage) *ImageService {

	return &ImageService{storage: storage}
}
func (s *ImageService) UploadFromFolder(ctx context.Context, folder string) error {
	files, err := os.ReadDir(folder)
	if err != nil {

		return fmt.Errorf("error read: %w", err)
	}
	for _, file := range files {
		path := folder + "/" + file.Name()
		f, err := os.Open(path)
		if err != nil {

			return fmt.Errorf("error wz file on path: %w", err)
		}
		img, err := jpeg.Decode(f)
		if err != nil {

			return fmt.Errorf("cant decode: %w", err)
		}
		var buf bytes.Buffer
		err = jpeg.Encode(&buf, img, &jpeg.Options{
			Quality: 60,
		})
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		size := int64(buf.Len())
		if err := s.storage.Upload(ctx, file.Name(), &buf, size); err != nil {

			return fmt.Errorf("upload failed: %w", err)
		}
		f.Close()
	}

	return nil
}
func (s *ImageService) DownloadToFolder(ctx context.Context, folder string) error {
	names, err := s.storage.List(ctx)
	if err != nil {

		return fmt.Errorf("cant list images: %w", err)
	}
	for _, name := range names {
		reader, err := s.storage.Download(ctx, name)
		if err != nil {
			return fmt.Errorf("error downloading: %w", err)
		}
		path := folder + "/" + name
		f, err := os.Create(path)
		if err != nil {

			return fmt.Errorf("error create file: %w", err)
		}
		_, err = io.Copy(f, reader)
		if err!=nil{

			return fmt.Errorf("copy failed: %w", err)
		}
		f.Close()
	}

	return nil
}
