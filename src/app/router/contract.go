// Package router
package router

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	pubsubrouter "github.com/sofyan48/pubsub-router"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
)

// httpHandlerFunc is a contract http handler for router
type httpHandlerFunc func(request *http.Request, svc contract.UseCase) appctx.Response

// Router is a contract router and must implement this interface
type Router interface {
	Route(ctx context.Context) *mux.Router
	TaskRoutes(ctx context.Context)
	KafkaProcessor(ctx context.Context)
	EventRouter() *pubsubrouter.Router
}
