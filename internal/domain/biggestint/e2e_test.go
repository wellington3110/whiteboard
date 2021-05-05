package biggestint_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wellington3110/whiteboard/internal/domain"
	"github.com/wellington3110/whiteboard/internal/domain/biggestint"
	"github.com/wellington3110/whiteboard/pkg/bootstraptest"
	"net/http"
	"testing"
)

const resource = "/biggest-int"

func TestFindBiggestInt_Success(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{{
		name:     "must find the biggest int in an ordered list",
		input:    "1,2,3,4",
		expected: 4,
	}, {
		name:     "must find the biggest int in an non ordered list",
		input:    "10, 2,-1,4,300",
		expected: 300,
	}, {
		name:     "must find the biggest int in an list containing just one item",
		input:    "10",
		expected: 10,
	}, {
		name:     "must trim all values",
		input:    "10, 2  ,  4,  12  ",
		expected: 12,
	}}
	server, client := bootstraptest.New(domain.Bootstrap())
	defer server.Close()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.NoError(t, client.Get(resource).
				AddQuery("ints", test.input).
				Expect(t).
				Status(http.StatusOK).
				JSON(map[string]int{
					"biggestInt": test.expected,
				}).
				Done())
		})
	}
}

func TestFindBiggestInt_Error(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expected   error
		statusCode int
	}{{
		name:       "must return an error when the list contains a non int value",
		input:      "1,2,3,a",
		expected:   biggestint.NewMustContainOnlyIntsError("1,2,3,a"),
		statusCode: http.StatusBadRequest,
	}, {
		name:       "must return an error when the list is empty",
		input:      "",
		expected:   biggestint.NewIntsMustNotBeEmptyError(),
		statusCode: http.StatusBadRequest,
	}}
	server, client := bootstraptest.New(domain.Bootstrap())
	defer server.Close()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.NoError(t, client.Get(resource).
				AddQuery("ints", test.input).
				Expect(t).
				Status(test.statusCode).
				JSON(test.expected).
				Done())
		})
	}
}
