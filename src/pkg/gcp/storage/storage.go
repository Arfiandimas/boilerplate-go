package storage

import (
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	cstore "cloud.google.com/go/storage"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/gcp"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

type stores struct {
	client *cstore.Client
}

func NewStorage(ctx context.Context, gcpA gcp.Contract) Storage {
	option := gcpA.Option()
	client, err := cstore.NewClient(ctx, option...)
	if err != nil {
		logger.Error("Client Error")
	}
	return &stores{
		client: client,
	}
}

func (s *stores) Save(ctx context.Context, file multipart.File, key string, contentType string) (uint64, error) {
	client := s.client.Bucket(os.Getenv("GOOGLE_BUCKET"))
	writer := client.Object(key).NewWriter(ctx)
	writer.ContentType = contentType
	result, err := io.Copy(writer, file)
	if err != nil {
		writer.Close()
		return 0, err
	}
	writer.Close()
	return uint64(result), nil
}

func (s *stores) Read(ctx context.Context, key string) ([]byte, error) {
	client := s.client.Bucket(os.Getenv("GOOGLE_BUCKET"))
	reader, err := client.Object(key).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}
