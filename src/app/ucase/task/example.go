package task

import (
	"context"
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/repositories"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/healthchecks"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

type expiredTX struct {
	logField []logger.Field
	repo     repositories.Example
	health   healthchecks.HealthchecksIO
}

func NewExpiredTX(repo repositories.Example, health healthchecks.HealthchecksIO) contract.TaskBackground {
	return &expiredTX{
		repo:   repo,
		health: health,
		logField: []logger.Field{
			logger.EventName("task_background"),
		},
	}
}

func (e *expiredTX) Run(ctx context.Context, t time.Time, done chan bool) error {
	return nil
}
