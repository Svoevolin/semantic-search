package httpclient

import (
	"net/http"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/delivery/http/middlewares"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func LoggingWrapper(logger logger.Logger) RoundTripperDecorator {
	const op = "httpclient.LoggingWrapper"
	return func(next http.RoundTripper) http.RoundTripper {
		return roundTripperFunc(func(r *http.Request) (*http.Response, error) {
			logger.InfoContext(r.Context(), op, "Outgoing request", "method", r.Method, "url", r.URL)

			resp, err := next.RoundTrip(r)
			if err != nil {
				logger.ErrorContext(r.Context(), op, "Request failed", "error", err.Error())
			} else {
				logger.InfoContext(r.Context(), op, "Response received", "status", resp.StatusCode)
			}

			return resp, err
		})
	}
}

func RequestIDWrapper() RoundTripperDecorator {
	return func(next http.RoundTripper) http.RoundTripper {
		return roundTripperFunc(func(r *http.Request) (*http.Response, error) {
			if requestID, ok := r.Context().Value(middlewares.HeaderRequestID).(string); ok && requestID != "" {
				r.Header.Set(string(middlewares.HeaderRequestID), requestID)
			}
			return next.RoundTrip(r)
		})
	}
}
