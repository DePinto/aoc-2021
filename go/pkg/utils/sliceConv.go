package utils

import (
	"strconv"
)

func NewIntSliceFromStringSlice(xs []string) ([]int, error) {
	xi := []int{}

	for _, s := range xs {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		xi = append(xi, i)
	}

	return xi, nil
}