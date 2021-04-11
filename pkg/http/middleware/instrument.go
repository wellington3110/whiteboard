package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) StatusCodeAsString() string {
	return strconv.Itoa(rw.statusCode)
}

var (
	requestTags        = []string{"path", "method", "code"}
	totalRequestsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		}, requestTags,
	)
	httpDurationHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, requestTags)
)

//nolint:gochecknoinits
func init() {
	prometheus.MustRegister(totalRequestsCount)
	prometheus.MustRegister(httpDurationHist)
}

func Instrument(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestInitTime := time.Now()
		rw := newResponseWriter(w)
		next.ServeHTTP(w, r)
		routePattern := chi.RouteContext(r.Context()).RoutePattern()
		if shouldSkipInstrumentation(routePattern) {
			return
		}
		tags := []string{routePattern, r.Method, rw.StatusCodeAsString()}
		totalRequestsCount.WithLabelValues(tags...).Inc()
		httpDurationHist.WithLabelValues(tags...).Observe(time.Since(requestInitTime).Seconds())
	})
}

func shouldSkipInstrumentation(routePattern string) bool {
	return routePattern == "" || strings.Contains(routePattern, "/metrics")
}
