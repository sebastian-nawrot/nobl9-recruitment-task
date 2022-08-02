package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/sync/errgroup"
)

type Results struct {
	Data              []int `json:"data"`
	StandardDeviation int   `json:"stddev"`
}

func fetchResults(requests int, length int) ([]Results, error) {
	results := make(chan []int, length)

	// Run workers and fetch results
	errorGroup := new(errgroup.Group)
	for i := 0; i < requests; i++ {
		errorGroup.Go(func() error {
			integers, err := makeRequest(length)
			if err == nil {
				results <- integers
			}
			return err
		})
	}

	// Check if error occurred in any of requests
	if err := errorGroup.Wait(); err != nil {
		return nil, err
	}

	accumulated := []int{}
	processedResults := []Results{}
	for i := 0; i < requests; i++ {
		integers := <-results
		processedResults = append(processedResults, Results{
			Data:              integers,
			StandardDeviation: standardDeviation(integers),
		})

		accumulated = append(accumulated, integers...)
	}

	// Standard deviation of sum of all sets
	processedResults = append(processedResults, Results{
		Data:              accumulated,
		StandardDeviation: standardDeviation(accumulated),
	})

	return processedResults, nil
}

func RandomMeanEndpoint(w http.ResponseWriter, r *http.Request) {
	errors := ErrorWrapper{}

	// Request method validation
	if r.Method != "GET" {
		errors.AppendError(APIError{
			Status:  http.StatusMethodNotAllowed,
			Message: fmt.Sprintf("Request method '%s' not supported", r.Method),
		})

		errors.Write(w, http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()

	// Request parameters validation
	requiredParameters := []string{"requests", "length"}
	for _, parameter := range requiredParameters {
		if !query.Has(parameter) {
			errors.AppendError(APIError{
				Status:    http.StatusBadRequest,
				Message:   fmt.Sprintf("Missing '%s' parameter", parameter),
				Parameter: parameter,
			})
		} else {
			value := query.Get(parameter)
			_, err := strconv.Atoi(value)
			if err != nil {
				errors.AppendError(APIError{
					Status:    http.StatusBadRequest,
					Message:   fmt.Sprintf("Invalid '%s' parameter format, expected number", parameter),
					Parameter: parameter,
				})
			}
		}
	}

	if len(errors.Errors) > 0 {
		errors.Write(w, http.StatusBadRequest)
		return
	}

	// Endpoint logic
	requests, _ := strconv.Atoi(query.Get("requests"))
	length, _ := strconv.Atoi(query.Get("length"))

	processedResults, err := fetchResults(requests, length)
	if err != nil {
		errors.AppendError(APIError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		errors.Write(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(processedResults)
}
