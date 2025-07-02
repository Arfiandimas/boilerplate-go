package router

import (
	"context"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/repositories"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/task"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/healthchecks"
)

func (rtr *router) TaskRoutes(ctx context.Context) {
	healthcheckIO := healthchecks.NewHeatlhChecksIO()
	examplerepo := repositories.NewExample(rtr.dbAdapter)
	expiredTask := task.NewExpiredTX(examplerepo, healthcheckIO)
	rtr.task(ctx, 900, expiredTask)
}
