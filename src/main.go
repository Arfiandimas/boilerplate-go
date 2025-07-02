// Package main
package main

import (
	"github.com/Arfiandimas/kaj-rest-engine-go/src/cmd"
)

// @title           {{PROJECT_NAME}} API
// @version         1.0
// @description     {{PROJECT_NAME}} API server for interact the app.
// @termsOfService  http://swagger.io/terms/

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					add your toke with this format | bearer (access_token) |
func main() {
	cmd.Start()
}
