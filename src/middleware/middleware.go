// Package middleware
package middleware

import (
	"net/http"
	"os"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/bootstrap"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/util"
)

// MiddlewareFunc is contract for middleware and must implement this type for http if need middleware http request
type MiddlewareFunc func(r *http.Request) string

// FilterFunc is a iterator resolver in each middleware registered
func FilterFunc(r *http.Request, mfs []MiddlewareFunc) string {
	for _, mf := range mfs {
		if status := mf(r); status != consts.MiddlewarePassed {
			return status
		}
	}

	return consts.MiddlewarePassed
}

func Oauth2MemberService(r *http.Request) string {
	if util.Environtment() == "loc" {
		return consts.MiddlewarePassed
	}
	client := bootstrap.RegistryOauthService()
	key, err := client.GetPublicKey()
	if err != nil {
		logger.Warn(`get public`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	rawToken := r.Header.Get("Authorization")
	_, token, err := util.ParseAccessToken(rawToken)
	if err != nil {
		logger.Warn(`parse token`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	tokenIssue, err := client.JWTParse(token, key)
	if err != nil {
		logger.Warn(`Error token`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}

	if !tokenIssue.Valid {
		logger.Warn(`token not valid`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	// comment if you dont need
	_, err = client.GetUserClaims(token)
	if err != nil {
		logger.Warn(`token not valid`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	if util.StringToBool(os.Getenv("OAUTH_TOKEN_VALIDATOR")) {
		err = client.Validator(rawToken, "member")
		if err != nil {
			return consts.AuthorizationFailure
		}
	}
	return consts.MiddlewarePassed
}

func Oauth2UserService(r *http.Request) string {
	if util.Environtment() == "loc" {
		return consts.MiddlewarePassed
	}
	client := bootstrap.RegistryOauthService()
	key, err := client.GetPublicKey()
	if err != nil {
		logger.Warn(`get public`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	rawToken := r.Header.Get("Authorization")
	_, token, err := util.ParseAccessToken(rawToken)
	if err != nil {
		logger.Warn(`parse token`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	tokenIssue, err := client.JWTParse(token, key)
	if err != nil {
		logger.Warn(`Error token`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}

	if !tokenIssue.Valid {
		logger.Warn(`token not valid`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	// comment if you dont need
	_, err = client.GetUserClaims(token)
	if err != nil {
		logger.Warn(`token not valid`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	if util.StringToBool(os.Getenv("OAUTH_TOKEN_VALIDATOR")) {
		err = client.Validator(rawToken, "user")
		if err != nil {
			return consts.AuthorizationFailure
		}
	}
	return consts.MiddlewarePassed
}

func Oauth2Service(r *http.Request) string {
	if util.Environtment() == "loc" {
		return consts.MiddlewarePassed
	}
	client := bootstrap.RegistryOauthService()
	key, err := client.GetPublicKey()
	if err != nil {
		logger.Warn(`get public`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	rawToken := r.Header.Get("Authorization")
	_, token, err := util.ParseAccessToken(rawToken)
	if err != nil {
		logger.Warn(`parse token`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	tokenIssue, err := client.JWTParse(token, key)
	if err != nil {
		logger.Warn(`Error token`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	if !tokenIssue.Valid {
		logger.Warn(`token not valid`,
			logger.SetField("error", err.Error()),
		)
		return consts.AuthorizationFailure
	}
	return consts.MiddlewarePassed
}
