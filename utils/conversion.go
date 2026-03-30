package utils

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {

	floats := make([]float64, len(strings))

	for stringIdx, stringVal := range strings {
		floatValue, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Failed to convert float")
		}
		floats[stringIdx] = floatValue
	}

	return floats, nil
}
