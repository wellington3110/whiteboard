package loggermiddleware_test

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/wellington3110/whiteboard/pkg/logger/loggermiddleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMiddlewareContext_WithRequestId(t *testing.T) {
	handlerCalled := false
	requestIDValue := "a"
	expectedMsg := "something"
	expectedRequestID := fmt.Sprintf(`"requestId":"%v"`, requestIDValue)
	req := createRequestWith(requestIDValue)
	handler := loggermiddleware.ContextApplying(loggermiddleware.RequestIDApply)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			outputLog := logMessageAndGetOutput(r, expectedMsg)
			require.Contains(t, outputLog, expectedRequestID)
			require.Contains(t, outputLog, expectedMsg)
			handlerCalled = true
		}))
	handler.ServeHTTP(httptest.NewRecorder(), req)
	require.True(t, handlerCalled)
}

func createRequestWith(requestIDValue string) *http.Request {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), middleware.RequestIDKey, requestIDValue))
	return req
}

func logMessageAndGetOutput(r *http.Request, expectedMsg string) string {
	writer := &strings.Builder{}
	logger := log.Ctx(r.Context()).Output(writer)
	logger.Info().Msg(expectedMsg)
	outputLog := writer.String()
	return outputLog
}
