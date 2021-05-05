package apierror_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/wellington3110/whiteboard/pkg/apierror"
	"testing"
)

func TestGetStatusCodeOr(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		fallback int
		expected int
	}{{
		name:     "should return status code from the error",
		err:      &apierror.APIError{StatusCode: 400},
		expected: 400,
	}, {
		name:     "should use fallback",
		err:      errors.New("some error"),
		fallback: 500,
		expected: 500,
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			statusCode := apierror.GetStatusCodeOr(test.err, test.fallback)
			require.Equal(t, test.expected, statusCode)
		})
	}
}
