package ucase

import (
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"
)

type healthCheck struct {
}

func NewHealthCheck() contract.UseCase {
	return &healthCheck{}
}

// @Summary Health Check Status
// @Description testing Check service health from server environment
// @Tags Healthcheck
// @Accept  json
// @Produce  json
// @Success 200 {object} appctx.Response
// @Failure 400 {object} appctx.Response
// @Failure 404 {object} appctx.Response
// @Failure 500 {object} appctx.Response
// @Router /in/health [get]
func (u *healthCheck) Serve(*appctx.Data) appctx.Response {
	return appctx.Response{
		Name:    consts.Success,
		Message: "ok",
	}
}
