package middleware

import (
	"github.com/alexliesenfeld/health"
)

// FullDetailsOnQueryParam is a middleware that removes check details (such as service names, error messages, etc.)
// from the HTTP response unless the request contained a query parameter named like argument 'queryParamName'. If
// a query parameter is not present in the HTTP request, the response will only contain the aggregated health status.
func FullDetailsOnQueryParam(queryParamName string) health.Middleware {
	_ = "STUB: not implemented"
	return *new(health.Middleware)
}
