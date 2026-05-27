package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	log "github.com/sirupsen/logrus"
)

// This example shows how to add check interceptors and handler middleware for pre- and post-processing.
// Both, interceptors and middleware allow to create re-usable functions, mostly used for cross-cutting
// functionality, such as logging, tracing, authentication, caching, etc.
func main() {

	// Create a new Checker
	checker := health.NewChecker(
		// A simple successFunc to see if a fake file system up.
		health.WithCheck(health.Check{
			Name:         "filesystem",
			Timeout:      2 * time.Second, // A successFunc specific timeout.
			Interceptors: []health.Interceptor{createCheckLogger, logCheck},
			Check: func(ctx context.Context) error {
				return fmt.Errorf("this is a check error") // example error
			},
		}),
	)

	handler := health.NewHandler(checker, health.WithMiddleware(createRequestLogger, logRequest))

	// We Create a new http.Handler that provides health successFunc information
	// serialized as a JSON string via HTTP.
	http.Handle("/health", handler)
	http.ListenAndServe(":3000", nil)
}

func createCheckLogger(next health.InterceptorFunc) health.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(health.InterceptorFunc)
}

func logCheck(next health.InterceptorFunc) health.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(health.InterceptorFunc)
}

func createRequestLogger(next health.MiddlewareFunc) health.MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(health.MiddlewareFunc)
}

func logRequest(next health.MiddlewareFunc) health.MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(health.MiddlewareFunc)
}

func setLogger(ctx context.Context, logger *log.Entry) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getLogger(ctx context.Context) *log.Entry { _ = "STUB: not implemented"; return nil }
