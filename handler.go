package health

import (
	"net/http"
)

type (
	HandlerConfig struct {
		statusCodeUp   int
		statusCodeDown int
		middleware     []Middleware
		resultWriter   ResultWriter
	}

	// Middleware is factory function that allows creating new instances of
	// a MiddlewareFunc. A MiddlewareFunc is expected to forward the function
	// call to the next MiddlewareFunc (passed in parameter 'next').
	// This way, a chain of interceptors is constructed that will eventually
	// invoke of the Checker.Check function. Each interceptor must therefore
	// invoke the 'next' interceptor. If the 'next' MiddlewareFunc is not called,
	// Checker.Check will never be executed.
	Middleware func(next MiddlewareFunc) MiddlewareFunc

	// MiddlewareFunc is a middleware for a health Handler (see NewHandler).
	// It is invoked each time an HTTP request is processed.
	MiddlewareFunc func(r *http.Request) CheckerResult

	// ResultWriter enabled a Handler (see NewHandler) to write the CheckerResult
	// to an http.ResponseWriter in a specific format. For example, the
	// JSONResultWriter writes the result in JSON format into the response body).
	ResultWriter interface {
		// Write writes a CheckerResult into a http.ResponseWriter in a format
		// that the ResultWriter supports (such as XML, JSON, etc.).
		// A ResultWriter is expected to write at least the following information into the http.ResponseWriter:
		// (1) A MIME type header (e.g., "Content-Type" : "application/json"),
		// (2) the HTTP status code that is passed in parameter statusCode (this is necessary due to ordering constraints
		// when writing into a http.ResponseWriter (see https://github.com/alexliesenfeld/health/issues/9), and
		// (3) the response body in the format that the ResultWriter supports.
		Write(result *CheckerResult, statusCode int, w http.ResponseWriter, r *http.Request) error
	}

	// JSONResultWriter writes a CheckerResult in JSON format into an
	// http.ResponseWriter. This ResultWriter is set by default.
	JSONResultWriter struct{}
)

// Write implements ResultWriter.Write.
func (rw *JSONResultWriter) Write(result *CheckerResult, statusCode int, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// NewJSONResultWriter creates a new instance of a JSONResultWriter.
func NewJSONResultWriter() *JSONResultWriter { _ = "STUB: not implemented"; return nil }

// NewHandler creates a new health check http.Handler.
func NewHandler(checker Checker, options ...HandlerOption) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Do the check (with configured middleware)

// Write HTTP response

//nolint:errcheck

func disableResponseCache(w http.ResponseWriter) {
	_ = "STUB: not implemented"
	// Avoid caching: https://www.ibm.com/garage/method/practices/manage/health-check-apis/
	return
}

func mapHTTPStatusCode(status AvailabilityStatus, statusCodeUp int, statusCodeDown int) int {
	_ = "STUB: not implemented"
	return 0
}

func createConfig(options []HandlerOption) HandlerConfig {
	_ = "STUB: not implemented"
	return *new(HandlerConfig)
}

func withMiddleware(interceptors []Middleware, target MiddlewareFunc) MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(MiddlewareFunc)
}
