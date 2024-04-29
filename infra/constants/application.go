package constants

import "errors"

const (
	DebugLog = "debug"
	InfoLog  = "info"
	ErrorLog = "error"
	WarnLog  = "warn"
)

type ExternalAPIError error

var (
	ErrorNoBaseURL ExternalAPIError = errors.New("base url is required")
	ErrorNoToken   ExternalAPIError = errors.New("bearer token is required")
)
