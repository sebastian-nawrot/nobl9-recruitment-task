package api

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Parameter string `json:"parameter,omitempty"`
}

type ErrorWrapper struct {
	Errors []APIError `json:"errors"`
}

func (ew *ErrorWrapper) AppendError(error APIError) {
	if error.Error == "" {
		error.Error = http.StatusText(error.Status)
	}

	ew.Errors = append(ew.Errors, error)
}

func (ew *ErrorWrapper) Write(w http.ResponseWriter, status int) {
	buf, _ := json.MarshalIndent(ew, "", " ")
	http.Error(w, string(buf), status)
}
