package kafkaCase

import (
	"os"
	"time"

	"github.com/kiriminaja/kaj-rest-engine-go/src/app/appctx"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/presentations"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/kiriminaja/kaj-rest-engine-go/src/consts"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/kafka"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

type subscribePublish struct {
	logField []logger.Field
	producer kafka.Producer
}

func NewPublish(producer kafka.Producer) contract.UseCase {
	return &subscribePublish{
		producer: producer,
		logField: []logger.Field{
			logger.EventName("kafka:producer"),
		},
	}
}

// @Summary Subscribe Publish
// @Description Example case for creating APIs
// @Tags Subscribe
// @Accept  json
// @Produce  json
// @Param data body presentations.PublishMessage true "Example Case Payload"
// @Success 200 {object} appctx.Response
// @Failure 400 {object} appctx.Response
// @Failure 404 {object} appctx.Response
// @Failure 500 {object} appctx.Response
// @Router /ex/v1/example/create [post]
func (u *subscribePublish) Serve(data *appctx.Data) appctx.Response {
	req := &presentations.PublishMessageRequest{}
	e := data.Cast(req)
	u.logField = append(u.logField, logger.Any("request", req))
	if e != nil {
		logger.Error(logger.SetMessageFormat("Parsing body request error: %s", e.Error()), u.logField...)
		return appctx.Response{
			Name:    consts.ValidationFailure,
			Message: e.Error(),
		}
	}

	now := time.Now()
	e = u.producer.Publish(data.Request.Context(), &kafka.MessageContext{
		Value: &kafka.BodyStateful{
			//wajib
			Body: req,
			//opsional
			Message: "Messages",
			Source: &kafka.SourceData{
				Service: os.Getenv("APP_NAME"),
			},
			// no need if you send request
			// Error: "",
		},
		Topic:     req.Topic,
		Partition: req.Partition,
		TimeStamp: now,
	})
	if e != nil {
		logger.Error(logger.SetMessageFormat("Error %v", e.Error()), u.logField...)
		return appctx.Response{
			Name: consts.InternalFailure,
		}
	}
	return appctx.Response{
		Name: consts.Success,
	}
}
