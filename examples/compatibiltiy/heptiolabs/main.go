package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/heptiolabs/healthcheck"
)

func main() {
	http.Handle("/health", health.NewHandler(
		health.NewChecker(
			health.WithCheck(health.Check{
				Name: "google",
				Check: func(ctx context.Context) error {
					deadline, _ := ctx.Deadline()
					timeout := time.Since(deadline)
					return healthcheck.HTTPGetCheck("https://www.google.com", timeout)()
				},
			}),
		)))
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
