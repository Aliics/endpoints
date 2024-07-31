# endpoints

Structured endpoints in Go.

The HTTP library in Go is powerful and extendable, and I think it allows us to easily extend it into structured
endpoints to more easily separate our logic and write tests with less setup.

# Example

This will handle our JSON (de)serialization, errors, status codes, timeouts, etc. and it's easily configured by
implementing a few interfaces.

```go
package main

import (
	"fmt"
	"github.com/aliics/endpoints"
	"log/slog"
	"net/http"
)

type testInput struct {
	Name string `json:"name"`
}

type testOutput struct {
	Result string `json:"result"`
}

type helloWorldEndpoint struct{}

func (e helloWorldEndpoint) EndpointPattern() string { return "POST /hello" }

func (e helloWorldEndpoint) Handle(in testInput) (*testOutput, error) {
	return &testOutput{fmt.Sprintf("Hello, %s!", in.Name)}, nil
}

func main() {
	mux := endpoints.NewEndpointsMux(helloWorldEndpoint{})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("http server failure", "err", err)
	}
}
```
