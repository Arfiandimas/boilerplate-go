package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/cmd/app"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/cmd/migration"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

// Start handler registering service command
func Start() {

	rootCmd := &cobra.Command{}
	logger.SetJSONFormatter()
	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	err := godotenv.Load()
	if err != nil {
		logger.Info(logger.SetMessageFormat("Environment load error: %s", err.Error()))
	}

	http := &cobra.Command{
		Use:   "http",
		Short: "Run http server",
	}
	http.AddCommand(app.Serve(ctx))

	cmd := []*cobra.Command{
		http,
		migration.Start(),
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
