package bootstrap

import (
	"context"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/gcp"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/gcp/pubsub"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/gcp/storage"
)

func GCPClient() gcp.Contract {
	return gcp.New()
}

func Storage(ctx context.Context, g gcp.Contract) storage.Storage {
	return storage.NewStorage(ctx, g)
}

func PubSub(ctx context.Context, g gcp.Contract) pubsub.PubSub {
	return pubsub.NewPubSub(ctx, g)
}
