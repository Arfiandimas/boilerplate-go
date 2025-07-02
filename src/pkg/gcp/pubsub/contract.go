package pubsub

import (
	"context"

	cpubsub "cloud.google.com/go/pubsub"
)

type PubSub interface {
	Client(ctx context.Context) *cpubsub.Client
	CreateTopic(ctx context.Context, topicID string) (*cpubsub.Topic, error)
	CreateSubscription(ctx context.Context, subID string, topic *cpubsub.Topic) (*cpubsub.Subscription, error)
	Publish(ctx context.Context, topic, path, message string) (string, error)
	PublishWithAttribute(ctx context.Context, topic, message string, attributes map[string]string) (string, error)
	DeleteSubscription(ctx context.Context, subID string) error
	DeleteTopic(ctx context.Context, topicID string) error
}
