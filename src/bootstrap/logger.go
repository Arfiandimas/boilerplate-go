// Package bootstrap
package bootstrap

import (
	"os"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/util"
)

// RegistryLogger initializer for logger
func RegistryLogger() {
	debug := util.StringToBool(os.Getenv("APP_DEBUG"))
	lc := logger.Config{
		URL:         os.Getenv("LOGGER_URL"),
		Environment: util.EnvironmentTransform(os.Getenv("APP_ENVIRONMENT")),
		Debug:       debug,
		Level:       os.Getenv("LOGGER_LEVEL"),
	}
	if lc.Environment != "loc" {
		h, e := logger.NewSentryHook(lc)

		if e != nil {
			logger.Fatal("log sentry failed to initialize Sentry")
		}
		logger.Setup(lc, h)
	}

}
