// Package handler
package handler

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/kafka"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

func KafkaProcessorHandler(ctx context.Context, consumer kafka.Consumer, producer kafka.Producer, topic []string, groupId string, handler contract.KafkaProcessor) {
	go func() error {
		if consumer == nil || producer == nil {
			logger.Error(logger.SetMessageFormat("Please attach lib in router handler"))
			return fmt.Errorf("Please attach lib in router handler")
		}
		consumer.Subscribe(&kafka.ConsumerContext{
			Handler: func(md *kafka.MessageDecoder) {
				data := &appctx.MessageDecoder{
					Body:      md.Body,
					Key:       md.Key,
					Message:   md.Message,
					Topic:     md.Topic,
					Partition: md.Partition,
					TimeStamp: md.TimeStamp,
					Offset:    md.Offset,
					Context:   ctx,
				}
				response := handler.Serve(data)
				response.Lang = os.Getenv("APP_DEFAULT_LANG")
				response.SetMessage()
				if kafka.KafkaDebugReport() {
					go sendDebugNotification(context.Background(), producer, &response, md.Source.Service, groupId)
				}

				lt := time.Since(md.TimeStamp)
				logger.AccessLog(md.Message,
					logger.Any("tag", "event-consumer"),
					logger.Any("group_consumer", groupId),
					logger.Any("response", response),
					logger.Any("topic", md.Topic),
					logger.Any("partition", md.Partition),
					logger.Any("offset", md.Offset),
					logger.Any("time", lt),
				)
			},
			GroupID: groupId,
			Topics:  topic,
			Context: ctx,
		})
		return nil
	}()
}

func sendDebugNotification(ctx context.Context, producer kafka.Producer, response *appctx.Response, service, consumerGroup string) {
	code := response.GetCode()
	bodyMessage := &kafka.BodyStateful{
		Body:    response.Data,
		Message: response.GetMessage(),
		Source: &kafka.SourceData{
			Service:       service,
			ConsumerGroup: consumerGroup,
		},
	}
	if response.Error != nil || !(code >= 200 && code <= 299) {
		if response.Error == nil {
			bodyMessage.Error = response.GetMessage()
		} else {
			bodyMessage.Error = response.Error.Error()
		}
	}
	topic := fmt.Sprintf("%s-%s-topic-debug", os.Getenv("APP_NAME"), os.Getenv("APP_ENVIRONMENT"))
	producer.Publish(ctx, &kafka.MessageContext{
		Value:     bodyMessage,
		Topic:     topic,
		TimeStamp: time.Now(),
	})
}
