package slice_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wellington3110/whiteboard/pkg/slice"
	"testing"
)

func TestFromStringToInt_Success(t *testing.T) {
	input := []string{"1", "3", "-1", "15", "1000"}
	expected := []int{1, 3, -1, 15, 1000}
	actual, err := slice.FromStringToInt(input)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestFromStringToInt_Error(t *testing.T) {
	input := []string{"1", "3", "a"}
	expectedErr := "strconv.Atoi: parsing \"a\": invalid syntax"
	actual, err := slice.FromStringToInt(input)
	require.Nil(t, actual)
	require.EqualError(t, err, expectedErr)
}

func TestTrimValues(t *testing.T) {
	input := []string{"1  ", " 2 ", "a    "}
	expected := []string{"1", "2", "a"}
	actual := slice.TrimValues(input)
	require.Equal(t, expected, actual)
}
