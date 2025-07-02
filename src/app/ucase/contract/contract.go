// Usecase Contract
package contract

import (
	"context"
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	pubsubrouter "github.com/sofyan48/pubsub-router"
)

// UseCase adalah jenis contract yang dapat digunakan untuk sebuah usecase HTTP
type UseCase interface {
	Serve(data *appctx.Data) appctx.Response
}

// MessageProcessor adalah jenis contract yang dapat digunakan untuk sebuah usecase MQ
type MessageProcessor interface {
	Serve(ctx context.Context) error
}

// TaskBackground adalah jenis contract yang dapat digunakan untuk sebuah usecase TASK
type TaskBackground interface {
	Run(ctx context.Context, t time.Time, done chan bool) error
}

type KafkaProcessor interface {
	Serve(data *appctx.MessageDecoder) appctx.Response
}

type EventProcessor interface {
	Serve(m *pubsubrouter.Message) error
}
