package collector

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Handler for test files
type MockHTTPHandler struct {
	ReturnJSON []byte
	Endpoint   string
}

// Used for testing
func (m *MockHTTPHandler) Get() (http.Response, error) {
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(m.ReturnJSON)),
	}
	return *response, nil
}
