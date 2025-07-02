package storage

import (
	"context"
	"mime/multipart"
)

type Storage interface {
	Save(ctx context.Context, file multipart.File, key string, contentType string) (uint64, error)
	Read(ctx context.Context, key string) ([]byte, error)
}
