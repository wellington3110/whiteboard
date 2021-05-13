package slice

import (
	"strconv"
	"strings"
)

func TrimValues(values []string) (result []string) {
	for _, s := range values {
		result = append(result, strings.TrimSpace(s))
	}
	return result
}

func FromStringToInt(values []string) ([]int, error) {
	var ints []int
	for _, s := range values {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, nil
}
