// Example List
package example

import (
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/appctx"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/repositories"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/kiriminaja/kaj-rest-engine-go/src/consts"

	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

type exampleList struct {
	repo     repositories.Example
	logField []logger.Field
}

// Usecase initializer for example list
func NewExampleList(repo repositories.Example) contract.UseCase {
	return &exampleList{
		repo: repo,
		logField: []logger.Field{
			logger.EventName("example:list"),
		},
	}
}

// @Summary Example List data
// @Description Example Get list data
// @Tags Example
// @Accept  json
// @Produce  json
// @Param limit query string true "10"
// @Param page query string true "0"
// @Success 200 {object} appctx.Response
// @Failure 400 {object} appctx.Response
// @Failure 404 {object} appctx.Response
// @Failure 500 {object} appctx.Response
// @Security Bearer
// @Router /ex/v1/example [get]
func (u *exampleList) Serve(data *appctx.Data) appctx.Response {
	pagination := &consts.Pagination{}
	err := data.Cast(pagination)
	u.logField = append(u.logField, logger.Any("pagination", pagination))
	if err != nil {
		logger.Warn(logger.SetMessageFormat("Parsing body pagination error: %s", err.Error()), u.logField...)
		return appctx.Response{
			Name:    consts.ValidationFailure,
			Message: err.Error(),
		}
	}
	listData, metaData, err := u.repo.Find(data.Request.Context(), pagination.Limit, pagination.Page)
	if err != nil {
		logger.Error(logger.SetMessageFormat("[merchant:list] %v", err.Error()), u.logField...)
		return appctx.Response{
			Name:    consts.InternalFailure,
			Message: err.Error(),
		}
	}

	return appctx.Response{
		Name:    consts.Success,
		Message: "List data",
		Data:    listData,
		Meta:    metaData,
	}
}
