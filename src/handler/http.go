// Package handler
package handler

import (
	"net/http"

	"github.com/kiriminaja/kaj-rest-engine-go/src/app/appctx"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/contract"
)

// HttpRequest handler func wrapper
func HttpRequest(request *http.Request, svc contract.UseCase) appctx.Response {
	data := &appctx.Data{
		Request: request,
	}

	return svc.Serve(data)
}
