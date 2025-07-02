// HTTP Usecase
//
//	Example Delete
package example

import (
	"github.com/gorilla/mux"
	"github.com/spf13/cast"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/appctx"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/repositories"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

type exampleDelete struct {
	repo     repositories.Example
	logField []logger.Field
}

// Usecase initializer for example delete
func NewExampleDelete(repo repositories.Example) contract.UseCase {
	return &exampleDelete{
		repo: repo,
		logField: []logger.Field{
			logger.EventName("example:delete"),
		},
	}
}

// @Summary Example Delete
// @Description Example Delete data
// @Tags Example
// @Accept  json
// @Produce  json
// @Param id   path      int  true  "Example ID"
// @Success 200 {object} appctx.Response
// @Failure 400 {object} appctx.Response
// @Failure 404 {object} appctx.Response
// @Failure 500 {object} appctx.Response
// @Security Bearer
// @Router /ex/v1/example/{id} [delete]
func (u *exampleDelete) Serve(data *appctx.Data) appctx.Response {
	id := mux.Vars(data.Request)["id"]
	u.logField = append(u.logField, logger.Any("id", id))
	e := u.repo.Delete(data.Request.Context(), cast.ToUint64(id))
	if e != nil {
		logger.Error(logger.SetMessageFormat("[example-delete] %v", e.Error()), u.logField...)
		return appctx.Response{
			Name: consts.InternalFailure,
		}
	}

	return appctx.Response{
		Name: consts.Success,
	}
}
