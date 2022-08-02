package api

import "testing"

func TestConvertStringSlice(t *testing.T) {
	for _, tt := range []struct {
		input  []string
		output []int
	}{
		{[]string{"1", "2"}, []int{1, 2}},
		{[]string{"9", "10"}, []int{9, 10}},
	} {
		numbers, err := convertStringSlice(tt.input)
		if err != nil {
			t.Errorf("Error occured during convertStringSlice execution: %s", err)
		}

		if len(numbers) != len(tt.output) {
			t.Errorf("Invalid length: %d != %d", len(numbers), len(tt.output))
		}

		for i, v := range numbers {
			if v != tt.output[i] {
				t.Errorf("Invalid element at %d index: %d != %d", i, v, tt.output[i])
			}
		}
	}
}

func TestStandardDeviation(t *testing.T) {
	for _, tt := range []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 1},
	} {
		sd := standardDeviation(tt.input)
		if sd != tt.output {
			t.Errorf("Invalid result: %d != %d", sd, tt.output)
		}
	}
}
