// Example Upsert
package example

import (
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/entity"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/presentations"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/repositories"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

type exampleCreate struct {
	repo     repositories.Example
	logField []logger.Field
}

// Usecase initializer for example create
func NewExampleCreate(repo repositories.Example) contract.UseCase {
	return &exampleCreate{
		repo: repo,
		logField: []logger.Field{
			logger.EventName("create_client"),
		},
	}
}

// @Summary Example Create
// @Description Example case for creating APIs
// @Tags Example
// @Accept  json
// @Produce  json
// @Param data body presentations.ExampleRequest true "Example Case Payload"
// @Success 200 {object} appctx.Response
// @Failure 400 {object} appctx.Response
// @Failure 404 {object} appctx.Response
// @Failure 500 {object} appctx.Response
// @Security Bearer
// @Router /ex/v1/example/create [post]
func (u *exampleCreate) Serve(data *appctx.Data) appctx.Response {
	req := presentations.ExampleRequest{}
	e := data.Cast(&req)
	u.logField = append(u.logField, logger.Any("reques", req))
	if e != nil {
		logger.Error(logger.SetMessageFormat("Parsing body request error: %s", e.Error()), u.logField...)
		return appctx.Response{
			Name:    consts.ValidationFailure,
			Message: e.Error(),
		}
	}
	now := time.Now()
	dataInsert := entity.Example{
		ID:        req.ID,
		Name:      req.Name,
		Address:   req.Address,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	id, e := u.repo.Upsert(data.Request.Context(), dataInsert)
	u.logField = append(u.logField, logger.Any("data", dataInsert))
	if e != nil {
		logger.Error(logger.SetMessageFormat("Error %v", e.Error()), u.logField...)
		return appctx.Response{
			Name: consts.InternalFailure,
		}
	}
	dataInsert.ID = id
	return appctx.Response{
		Name: consts.Created,
		Data: dataInsert,
	}
}
