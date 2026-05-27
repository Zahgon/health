package interceptors

import (
	"github.com/alexliesenfeld/health"
)

// BasicLogger is a basic logger that is mostly used to showcase this library.
func BasicLogger() health.Interceptor { _ = "STUB: not implemented"; return *new(health.Interceptor) }
