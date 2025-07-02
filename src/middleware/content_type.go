package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kiriminaja/kaj-rest-engine-go/src/consts"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

// ValidateContentType header
func ValidateContentType(r *http.Request) string {
	if ct := strings.ToLower(r.Header.Get(`Content-Type`)); ct != `application/json` {
		logger.Warn(fmt.Sprintf("[middleware] invalid content-type %s", ct))

		return consts.UnprocessAbleEntity
	}

	return consts.MiddlewarePassed
}
