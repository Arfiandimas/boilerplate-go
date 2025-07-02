package pubsub

import (
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	pubsubrouter "github.com/sofyan48/pubsub-router"
)

type defaultRoutes struct {
	logField []logger.Field
}

func NewDefaultRoutes() contract.EventProcessor {
	return &defaultRoutes{
		logField: []logger.Field{
			logger.EventName("no_routes_handler"),
		},
	}
}

// Serve implements contract.EventProcessor.
func (d *defaultRoutes) Serve(m *pubsubrouter.Message) error {
	d.logField = append(d.logField, logger.Any("data", string(m.Data)))
	logger.Warn(logger.SetMessageFormat("Receive no any match with routes: %s", m.Attribute["path"]), d.logField...)
	return nil
}
