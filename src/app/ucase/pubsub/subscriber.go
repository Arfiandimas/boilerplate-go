package pubsub

import (
	"fmt"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	pubsubrouter "github.com/sofyan48/pubsub-router"
)

type subscription struct {
	logField []logger.Field
}

func NewSubscriber() contract.EventProcessor {
	return &subscription{
		logField: []logger.Field{
			logger.EventName("subscriber"),
		},
	}
}

// Serve implements contract.EventProcessor.
func (*subscription) Serve(m *pubsubrouter.Message) error {
	fmt.Println(string(m.Data))
	return nil
}
