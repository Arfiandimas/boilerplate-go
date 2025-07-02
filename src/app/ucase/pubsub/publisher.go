package pubsub

import (
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/appctx"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/presentations"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/kiriminaja/kaj-rest-engine-go/src/consts"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/gcp/pubsub"
)

type publisher struct {
	client pubsub.PubSub
}

func NewPublisher(client pubsub.PubSub) contract.UseCase {
	return &publisher{
		client: client,
	}
}

// Serve implements contract.UseCase.
func (c *publisher) Serve(data *appctx.Data) appctx.Response {
	req := &presentations.PublishMessagePubsub{}
	err := data.Cast(req)
	if err != nil {
		return appctx.Response{
			Name:    consts.ValidationFailure,
			Message: err.Error(),
		}
	}

	sts, err := c.client.Publish(data.Request.Context(), req.Topic, "/test", req.Messages)
	if err != nil {
		return appctx.Response{
			Name:    consts.InternalFailure,
			Message: err.Error(),
		}
	}
	return appctx.Response{
		Name:    consts.Success,
		Message: "published",
		Data:    sts,
	}
}
