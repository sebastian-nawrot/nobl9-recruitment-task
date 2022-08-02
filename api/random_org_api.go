package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const endpoint = "https://www.random.org/integers/"

func makeRequest(length int) ([]int, error) {
	// Constraints defined in random.org API, check https://www.random.org/clients/http/
	if length <= 0 && length > 1000 {
		return nil, errors.New("parameter 'length' has to be within the range [1,1000]")
	}

	query := fmt.Sprintf("?num=%d&min=1&max=10&col=1&base=10&format=plain&rnd=new", length)

	response, err := http.Get(endpoint + query)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	trimmed_strings := strings.TrimSuffix(string(response_body), "\n")
	numbers := strings.Split(trimmed_strings, "\n")
	return convertStringSlice(numbers)
}
