// Package bootstrap
package bootstrap

import (
	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/msg"
)

// RegistryMessage setup transformation
func RegistryMessage() {
	err := msg.Setup("msg.yaml", consts.CONFIG_PATH)
	if err != nil {
		logger.Fatal(logger.SetMessageFormat("file message multi language load error %s", err.Error()))
	}

}
