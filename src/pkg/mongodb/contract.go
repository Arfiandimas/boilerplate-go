package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Adapter interface {
	SetCollection(name string, opts *options.CollectionOptions) *mongo.Collection
	Ping(ctx context.Context)
	Upsert(ctx context.Context, collection string, id uint64, data interface{}) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, collection string, id uint64) (*mongo.DeleteResult, error)
}

type Config struct {
	Username string
	Password string
	Port     string
	Host     string
	Timeout  int
	Name     string
}
