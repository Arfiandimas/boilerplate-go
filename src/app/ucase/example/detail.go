// Example Detail
package example

import (
	"github.com/gorilla/mux"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/appctx"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/repositories"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/kiriminaja/kaj-rest-engine-go/src/consts"

	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/util"
)

type exampleDetail struct {
	repo     repositories.Example
	logField []logger.Field
}

// Usecase initializer for example detail
func NewExampleDetail(repo repositories.Example) contract.UseCase {
	return &exampleDetail{
		repo: repo,
		logField: []logger.Field{
			logger.EventName("example:detail"),
		},
	}
}

// @Summary Example Detail
// @Description Example Get list data
// @Tags Example
// @Accept  json
// @Produce  json
// @Param id   path      int  true  "Example ID"
// @Success 200 {object} appctx.Response
// @Failure 400 {object} appctx.Response
// @Failure 404 {object} appctx.Response
// @Failure 500 {object} appctx.Response
// @Security Bearer
// @Router /ex/v1/example/{id} [get]
func (u *exampleDetail) Serve(data *appctx.Data) appctx.Response {
	id := mux.Vars(data.Request)["id"]
	detailData, err := u.repo.Fetch(data.Request.Context(), util.StrToUint64(id))
	u.logField = append(u.logField, logger.Any("id", id))
	if err != nil {
		logger.Error(logger.SetMessageFormat("[merchant:detail] %v", err.Error()), u.logField...)
		return appctx.Response{
			Name:    consts.InternalFailure,
			Message: err.Error(),
		}
	}

	return appctx.Response{
		Name:    consts.Success,
		Message: "List data",
		Data:    detailData,
	}
}
