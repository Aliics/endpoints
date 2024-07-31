package endpoints

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type testInput struct {
	Name string `json:"name"`
}

type testValidatedInput struct {
	RequiredName string `json:"requiredName"`
}

func (t testValidatedInput) Validate() error {
	if t.RequiredName == "" {
		return errors.New("requiredName is required... duh")
	}

	return nil
}

type testOutput struct {
	Result string `json:"result"`
}

type testEndpointWithBodyHandler struct{}

func (e testEndpointWithBodyHandler) EndpointPattern() string { return "POST /test" }

func (e testEndpointWithBodyHandler) Handle(in testInput) (*testOutput, error) {
	return &testOutput{fmt.Sprintf("Hello, %s!", in.Name)}, nil
}

type testEndpointWithBodyHandlerWithValidatedInput struct{}

func (e testEndpointWithBodyHandlerWithValidatedInput) EndpointPattern() string { return "POST /test" }

func (e testEndpointWithBodyHandlerWithValidatedInput) Handle(in testValidatedInput) (*testOutput, error) {
	return &testOutput{fmt.Sprintf("Hello, %s!", in.RequiredName)}, nil
}

type testEndpointWithErrOnlyHandler struct{}

func (e testEndpointWithErrOnlyHandler) EndpointPattern() string { return "POST /err-only-test" }

func (e testEndpointWithErrOnlyHandler) Handle(in testInput) error {
	return BadRequestError(fmt.Sprintf("failed with input: %v", in))
}

type testEndpointWithTimeoutHandler struct{}

func (e testEndpointWithTimeoutHandler) EndpointPattern() string { return "POST /timeout-test" }

func (e testEndpointWithTimeoutHandler) WithTimeout() time.Duration { return 10 * time.Millisecond }

func (e testEndpointWithTimeoutHandler) Handle(_ context.Context) error {
	time.Sleep(20 * time.Millisecond)
	return nil
}

func TestEndpointMuxWithBodyHandlerEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"name":"Alex"}`))
	rec := httptest.NewRecorder()

	NewEndpointsMux(testEndpointWithBodyHandler{}).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"result\":\"Hello, Alex!\"}\n", rec.Body.String())
}

func TestEndpointMuxWithBodyHandlerEndpointAndValidatedInput(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"requiredName":""}`))
	rec := httptest.NewRecorder()

	NewEndpointsMux(testEndpointWithBodyHandlerWithValidatedInput{}).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "requiredName is required... duh\n", rec.Body.String())
}

func TestEndpointMuxWithErrOnlyHandlerEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/err-only-test", strings.NewReader(`{"name":"oh noes"}`))
	rec := httptest.NewRecorder()

	NewEndpointsMux(testEndpointWithErrOnlyHandler{}).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "failed with input: {oh noes}\n", rec.Body.String())
}

func TestEndpointMuxWithTimeoutHandlerEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/timeout-test", strings.NewReader(`{"name":"oh noes"}`))
	rec := httptest.NewRecorder()

	NewEndpointsMux(testEndpointWithTimeoutHandler{}).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusRequestTimeout, rec.Code)
}
