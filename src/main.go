// Package main
package main

import (
	"github.com/kiriminaja/kaj-rest-engine-go/src/cmd"
)

// @title           {{PROJECT_NAME}} API
// @version         1.0
// @description     {{PROJECT_NAME}} API server for interact the app.
// @termsOfService  http://swagger.io/terms/

// @license.name  Kiriminaja License 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					add your toke with this format | bearer (access_token) |
func main() {
	cmd.Start()
}
