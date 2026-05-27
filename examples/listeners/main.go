package main

import (
	"context"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
)

// This example shows how this library can be used to listen for health status changes.
func main() {
	// Create a new Checker
	checker := health.NewChecker(

		// Configure a global timeout that will be applied to all checks.
		health.WithTimeout(10*time.Second),

		// A simple successFunc to see if a fake database connection is up.
		health.WithCheck(health.Check{
			Name:           "database",
			Timeout:        2 * time.Second, // A successFunc specific timeout.
			StatusListener: onComponentStatusChanged,
			Check: func(ctx context.Context) error {
				return nil // no error
			},
		}),

		// A simple successFunc to see if a fake file system up.
		health.WithCheck(health.Check{
			Name:           "filesystem",
			Timeout:        2 * time.Second, // A successFunc specific timeout.
			StatusListener: onComponentStatusChanged,
			Check: func(ctx context.Context) error {
				return nil // example error
			},
		}),

		// The following check will be executed periodically every 10 seconds.
		health.WithPeriodicCheck(5*time.Second, 10*time.Second, health.Check{
			Name:           "search-engine",
			StatusListener: onComponentStatusChanged,
			Check:          volatileFunc(),
		}),

		// This listener will be called whenever system health status changes (e.g., from "up" to "down").
		health.WithStatusListener(onSystemStatusChanged),
	)

	// We Create a new http.Handler that provides health successFunc information
	// serialized as a JSON string via HTTP.
	http.Handle("/health", health.NewHandler(checker))
	http.ListenAndServe(":3000", nil)
}

func onComponentStatusChanged(_ context.Context, name string, state health.CheckState) {
	_ = "STUB: not implemented"
	return
}

func onSystemStatusChanged(_ context.Context, state health.CheckerState) {
	_ = "STUB: not implemented"
	return
}

func volatileFunc() func(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// example error
