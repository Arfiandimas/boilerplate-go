package pubsub

import (
	"context"
	"os"
	"time"

	cpubsub "cloud.google.com/go/pubsub"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/gcp"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

type pubsubGCP struct {
	client *cpubsub.Client
}

func NewPubSub(ctx context.Context, gcpA gcp.Contract) PubSub {
	option := gcpA.Option()
	client, err := cpubsub.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"), option...)
	if err != nil {
		logger.Error("Client Error")
		defer client.Close()
	}
	return &pubsubGCP{
		client: client,
	}
}

func (p *pubsubGCP) Client(ctx context.Context) *cpubsub.Client {
	return p.client
}

func (p *pubsubGCP) CreateTopic(ctx context.Context, topicID string) (*cpubsub.Topic, error) {
	result, err := p.client.CreateTopic(ctx, topicID)
	if err != nil {
		logger.Error(logger.SetMessageFormat("Error %v", err.Error()))
		return nil, err
	}
	logger.Info("Creating and sycn topic")
	return result, nil
}

func (p *pubsubGCP) CreateSubscription(ctx context.Context, subID string, topic *cpubsub.Topic) (*cpubsub.Subscription, error) {
	result, err := p.client.CreateSubscription(ctx, subID, cpubsub.SubscriptionConfig{
		Topic:             topic,
		AckDeadline:       10 * time.Second,
		RetentionDuration: 24 * time.Hour,
		// DeadLetterPolicy: &cpubsub.DeadLetterPolicy{
		// 	DeadLetterTopic:     "DEAD-LETTER",
		// 	MaxDeliveryAttempts: 5,
		// },
	})
	if err != nil {
		logger.Error(logger.SetMessageFormat("Error %v", err.Error()))
		return nil, err
	}
	logger.Info("Creating and sycn subscription topic")
	return result, nil
}

func (p *pubsubGCP) DeleteSubscription(ctx context.Context, subID string) error {
	err := p.client.Subscription(subID).Delete(ctx)
	if err != nil {
		logger.Error(logger.SetMessageFormat("Error %v", err.Error()))
		return err
	}
	logger.Info("Delete subscription topic")
	return nil
}

func (p *pubsubGCP) DeleteTopic(ctx context.Context, topicID string) error {
	err := p.client.Topic(topicID).Delete(ctx)
	if err != nil {
		logger.Error(logger.SetMessageFormat("Error %v", err.Error()))
		return err
	}
	logger.Info("Delete topic")
	return nil
}

func (p *pubsubGCP) PublishWithAttribute(ctx context.Context, topic, message string, attributes map[string]string) (string, error) {
	cl := p.client.Topic(topic)
	result := cl.Publish(
		ctx,
		&cpubsub.Message{
			ID:          "",
			Data:        []byte(message),
			Attributes:  attributes,
			PublishTime: time.Now(),
		},
	)
	res, err := result.Get(ctx)
	if err != nil {
		logger.Error(logger.SetMessageFormat("Publish Error %v", err.Error()))
		return "", err
	}
	return res, nil
}

func (p *pubsubGCP) Publish(ctx context.Context, topic, path, message string) (string, error) {
	cl := p.client.Topic(topic)
	result := cl.Publish(
		ctx,
		&cpubsub.Message{
			Data:        []byte(message),
			PublishTime: time.Now(),
			Attributes: map[string]string{
				"path": path,
			},
		},
	)
	res, err := result.Get(ctx)
	if err != nil {
		logger.Error(logger.SetMessageFormat("Publish Error %v", err.Error()))
		return "", err
	}
	logger.Info(logger.SetMessageFormat("Publish Finish %v", res))
	return res, nil
}
