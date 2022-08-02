package api

import "testing"

func TestFetchResults(t *testing.T) {
	const expectedRequests = 1
	const expectedLength = 4
	processedResults, err := fetchResults(expectedRequests, expectedLength)
	if err != nil {
		t.Errorf("Error occured during fetchResults execution: %s", err)
	}

	if len(processedResults) != expectedRequests+1 {
		t.Errorf("Invalid length of resutls: %d != %d",
			len(processedResults), expectedLength)
	}

	// Test standard deviation of sum of all sets
	firstResult := processedResults[0]
	sumResult := processedResults[1]
	if firstResult.StandardDeviation != sumResult.StandardDeviation {
		t.Errorf("Invalid standard deviation of sum of all sets: %d != %d",
			firstResult.StandardDeviation, sumResult.StandardDeviation)
	}
}
