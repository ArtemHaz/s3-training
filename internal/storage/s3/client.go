package s3

import (
	"context"
	"fmt"
	"io"
	"s3/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	client *minio.Client
	bucket string
}

func NewClient(cfg *config.Config) (*Client, error) {
	minioClient, err := minio.New(cfg.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Secure: false,
	})
	if err != nil {

		return nil, fmt.Errorf("create minio client: %w", err)
	}

	return &Client{
		client: minioClient,
		bucket: cfg.S3Bucket,
	}, nil
}

func (c *Client) Upload(ctx context.Context, name string, reader io.Reader, size int64) error {
	_, err := c.client.PutObject(ctx, c.bucket, name, reader, size, minio.PutObjectOptions{})
	if err != nil {

		return fmt.Errorf("put object: %w", err)
	}

	return nil
}
func (c *Client) Download(ctx context.Context, name string) (io.Reader, error) {
	obj, err := c.client.GetObject(ctx, c.bucket, name, minio.GetObjectOptions{})
	if err != nil {

		return nil, fmt.Errorf("cant get obj: %w", err)
	}

	return obj, nil
}

func (c *Client) List(ctx context.Context) ([]string, error) {
	objects := c.client.ListObjects(ctx, c.bucket, minio.ListObjectsOptions{})
	var names []string
	for obj := range objects {
		if obj.Err != nil {
			return nil, obj.Err
		}
		names = append(names, obj.Key)
	}

	return names, nil
}
