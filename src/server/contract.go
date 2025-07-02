// Package router
package server

import (
	"context"
	"net/http"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	ucase "github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
)

// httpHandlerFunc abstraction for http handler
type httpHandlerFunc func(request *http.Request, svc ucase.UseCase) appctx.Response

// Server contract
type Server interface {
	Run(ctx context.Context, port string) error
	EventProcessor(ctx context.Context)
	Done()
}
