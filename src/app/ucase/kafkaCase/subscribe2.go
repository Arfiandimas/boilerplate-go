// MQ (sqs) subscriber example
package kafkaCase

import (
	"fmt"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/presentations"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/repositories"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

type subscribeExample struct {
	logField []logger.Field
	repo     repositories.Example
}

func NewSubscribe(repo repositories.Example) contract.KafkaProcessor {
	return &subscribeExample{
		repo: repo,
		logField: []logger.Field{
			logger.EventName("kafka:subscribe"),
		},
	}
}

// Serve implements contract.KafkaProcessor
func (s *subscribeExample) Serve(data *appctx.MessageDecoder) appctx.Response {
	fmt.Println("SUBS =======> 1")
	bodyData := presentations.PublishMessageRequest{}
	e := data.Cast(&bodyData)
	if e != nil {
		return appctx.Response{
			Name:    consts.ValidationFailure,
			Message: "Validation Failure",
		}
	}
	s.repo.Find(data.Context, 10, 1)
	return appctx.Response{
		Name: consts.Success,
		Data: bodyData,
	}
}
