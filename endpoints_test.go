package endpoints

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testInput struct {
	Name string `json:"name"`
}

type testOutput struct {
	Result string `json:"result"`
}

type testEndpointWithBodyHandler struct{}

func (t testEndpointWithBodyHandler) EndpointPattern() string {
	return "POST /test"
}

func (t testEndpointWithBodyHandler) Handle(in testInput) (*testOutput, error) {
	return &testOutput{fmt.Sprintf("Hello, %s!", in.Name)}, nil
}

func TestEndpointMuxWithBodyHandlerEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"name":"Alex"}`))
	rec := httptest.NewRecorder()

	NewEndpointsMux(testEndpointWithBodyHandler{}).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"result\":\"Hello, Alex!\"}\n", rec.Body.String())
}
