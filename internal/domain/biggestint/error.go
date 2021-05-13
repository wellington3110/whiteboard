package biggestint

import (
	"github.com/wellington3110/whiteboard/pkg/apierror"
	"net/http"
)

type (
	MustContainOnlyIntsError struct{ *apierror.APIError }
	IntsMustNotBeEmptyError  struct{ *apierror.APIError }
)

const (
	mustContainOnlyIntsErrorCode = "must_contain_only_ints_error"
	intsMustNotBeEmptyErrorCode  = "ints_must_not_be_empty_error"
)

func NewMustContainOnlyIntsError(input string) *MustContainOnlyIntsError {
	return &MustContainOnlyIntsError{&apierror.APIError{
		Code:       mustContainOnlyIntsErrorCode,
		Message:    "you must input only ints",
		StatusCode: http.StatusBadRequest,
		Context: map[string]interface{}{
			"input": input,
		},
	}}
}

func NewIntsMustNotBeEmptyError() *IntsMustNotBeEmptyError {
	return &IntsMustNotBeEmptyError{&apierror.APIError{
		Code:       intsMustNotBeEmptyErrorCode,
		Message:    "you must input a non-empty list of ints",
		StatusCode: http.StatusBadRequest,
	}}
}
