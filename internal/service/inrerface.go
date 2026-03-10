package service

import (
	"context"
	"io"
)

type ImageStorage interface {
	Upload(ctx context.Context, name string, reader io.Reader, size int64) error
	Download(ctx context.Context, name string) (io.Reader, error)
	List(ctx context.Context) ([]string, error)
}
