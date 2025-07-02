package handler

import (
	"context"
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase/contract"
)

func TaskHandler(sleep int64, task contract.TaskBackground, ctx context.Context) {
	go func(sleep int64, task contract.TaskBackground, ctx context.Context) {
		done := make(chan bool)
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case t := <-ticker.C:
				time.Sleep(time.Duration(sleep) * time.Second)
				task.Run(ctx, t, done)
			}
		}
	}(sleep, task, ctx)
}
