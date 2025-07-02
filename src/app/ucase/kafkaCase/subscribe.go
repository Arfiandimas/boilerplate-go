// MQ (sqs) subscriber example
package kafkaCase

import (
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/presentations"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/repositories"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

type subscribeExample2 struct {
	logField []logger.Field
	repo     repositories.Example
}

func NewSubscribe2(repo repositories.Example) contract.KafkaProcessor {
	return &subscribeExample2{
		repo: repo,
		logField: []logger.Field{
			logger.EventName("kafka:subscribe"),
		},
	}
}

// Serve implements contract.KafkaProcessor
func (s *subscribeExample2) Serve(data *appctx.MessageDecoder) appctx.Response {
	bodyData := presentations.PublishMessageRequest{}
	e := data.Cast(&bodyData)
	if e != nil {
		return appctx.Response{
			Name:    consts.ValidationFailure,
			Message: "Validation Failure",
		}
	}
	return appctx.Response{
		Name:    consts.Success,
		Message: "From message",
		Data:    bodyData,
	}
}
