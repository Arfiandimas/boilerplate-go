package app

import (
	"context"
	"os"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/server"
	"github.com/spf13/cobra"
)

func Serve(ctx context.Context) *cobra.Command {
	serve := &cobra.Command{}
	serve.Flags().StringP("port", "p", os.Getenv("SERVER_PORT"), "Add port")
	serve.Use = "serve"
	serve.Short = "Run Http Server"
	serve.Run = func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		serve := server.NewHTTPServer()
		// serve.EventProcessor(ctx)
		defer serve.Done()
		logger.Info(logger.SetMessageFormat("starting services... " + os.Getenv("SERVER_PORT")))
		if err := serve.Run(ctx, port); err != nil {
			logger.Warn(logger.SetMessageFormat("service stopped, err:%s", err.Error()))
		}
	}
	return serve
}
