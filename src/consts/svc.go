// Package consts
package consts

const (
	DATE_LAYOUT_DEFAULT = `2006-01-02 15:04:05`

	// ENV_PRODDUCTION name
	ENV_PRODDUCTION = `prod`
	// ENV_STAGING name
	ENV_STAGING = `stg`
	// ENV_DEVELOPMENT name
	ENV_DEVELOPMENT = `dev`

	//ServiceTypeHTTP marks the usecase service type for HTTP operation
	ServiceTypeHTTP = "http"
	//ServiceTypeMQ marks the usecasee service type for MQ Operation
	ServiceTypeConsumer = "consumer"

	// MiddlewarePassed response var for middleware http request passed
	MiddlewarePassed = `MIDDLEWARE_PASSED`

	// 2xx status code

	// Success response var for general success
	Success = "SUCCESS"
	// Created Success response fo general created
	Created = "CREATED"
	// Accepted Success response fo general accept action
	Accepted = "ACCEPTED"
	// end 2xx status code

	// 4xx status code

	// AuthenticationFailure response var
	AuthenticationFailure = `AUTHENTICATION_FAILURE`
	// AuthorizationFailure ...
	AuthorizationFailure = `UNAUTHORIZED_FAILURE`
	// SignatureFaulure response var
	SignatureFaulure = `SIGNATURE_FAILURE`
	// ValidationFailure response var for general validation error or not pass
	ValidationFailure = `VALIDATION_FAILURE`
	// ValidationInvalidEmail response var
	ValidationInvalidEmail = `VALIDATION_INVALID_EMAIL`
	// DataNotFound response var for general not found data in our system or third party services
	DataNotFound = `DATA_NOT_FOUND`
	// UnprocessAbleEntity response var for general wen we cannot continue process and can not retry
	UnprocessAbleEntity = `UNPROCESSABLE_ENTITY`
	// Forbidden response var for forbidden access
	Forbidden = `FORBIDDEN`
	// RequestTimeOut response var
	RequestTimeOut = `REQUEST_TIMEOUT`
	// 4xx status code

	// InternalFailure response var for internal server error or like something went wrong in system
	InternalFailure = `INTERNAL_FAILURE`
	// ServiceUnavailable response var when tah service is not available
	ServiceUnavailable = `SERVICE_UNAVAILABLE`
	// UnknownError response var for unknown error
	UnknownError = `UNKNOWN_ERROR`
)

type Pagination struct {
	Page  uint64 `json:"page" form:"page" validate:"min=1"`
	Limit uint64 `json:"limit" form:"limit" validate:"min=1"`
}
