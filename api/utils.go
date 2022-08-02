package api

import (
	"math"
	"strconv"
)

func convertStringSlice(data []string) ([]int, error) {
	integers := []int{}
	for i := range data {
		converted_value, err := strconv.Atoi(data[i])
		if err != nil {
			return nil, err
		}

		integers = append(integers, converted_value)
	}

	return integers, nil
}

func standardDeviation(data []int) int {
	sum := 0
	for i := range data {
		sum += data[i]
	}

	mean := sum / len(data)

	sd := 0
	for i := range data {
		sd += int(math.Pow(float64(data[i]-mean), 2))
	}

	return int(math.Sqrt(float64(sd / 10)))
}
