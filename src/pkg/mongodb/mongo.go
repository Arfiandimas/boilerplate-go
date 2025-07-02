package mongodb

import (
	"context"

	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoClient(cfg *Config) (Adapter, error) {
	x := &mongoDB{}
	session, db, err := CreateSession(cfg)
	if err != nil {
		return nil, err
	}
	x.client = session
	x.db = db
	return x, nil
}

func (m *mongoDB) SetCollection(name string, opts *options.CollectionOptions) *mongo.Collection {
	return m.db.Collection(name, opts)
}

func (m *mongoDB) Ping(ctx context.Context) {
	m.client.Ping(ctx, nil)
}

func (m *mongoDB) Upsert(ctx context.Context, collection string, id uint64, data interface{}) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}
	opts := options.Update().SetUpsert(true)
	return m.db.Collection(collection, options.Collection()).UpdateOne(ctx, filter, update, opts)
}

func (m *mongoDB) Delete(ctx context.Context, collection string, id uint64) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(util.ToString(id))
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	return m.db.Collection(collection).DeleteOne(ctx, filter)
}
