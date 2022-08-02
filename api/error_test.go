package api

import (
	"net/http"
	"testing"
)

func TestErrorCreation(t *testing.T) {
	errors := ErrorWrapper{}
	errors.AppendError(APIError{
		Status:  http.StatusNotFound,
		Message: "Couldn't find requested resource",
	})

	err := errors.Errors[0]
	if err.Status != http.StatusNotFound {
		t.Errorf("Invalid error Status: %d", err.Status)
	}

	if err.Error != http.StatusText(http.StatusNotFound) {
		t.Errorf("Invalid error Error: %s", err.Error)
	}

	if err.Message != "Couldn't find requested resource" {
		t.Errorf("Invalid error Message: %s", err.Message)
	}
}
