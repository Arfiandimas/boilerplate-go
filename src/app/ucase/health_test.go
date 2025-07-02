// Package ucase
package ucase

import (
	"testing"

	"github.com/kiriminaja/kaj-rest-engine-go/src/app/appctx"
	"github.com/kiriminaja/kaj-rest-engine-go/src/consts"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_Serve(t *testing.T) {
	svc := NewHealthCheck()

	t.Run("test health check", func(t *testing.T) {
		result := svc.Serve(&appctx.Data{})

		assert.Equal(t, appctx.Response{
			Name:    consts.Success,
			Message: "ok",
		}, result)
	})
}
