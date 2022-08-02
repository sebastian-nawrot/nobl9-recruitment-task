package api

import "testing"

func TestMakeRequest(t *testing.T) {
	const expectedLength = 8
	numbers, err := makeRequest(expectedLength)
	if err != nil {
		t.Errorf("Error occured during makeRequests execution: %s", err)
	}

	if len(numbers) != expectedLength {
		t.Errorf("Invalid length of resutls: %d != %d", len(numbers), expectedLength)
	}
}
