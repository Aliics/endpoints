package endpoints

import "time"

// EndpointHandler allows for the registration of an endpoint with a route.
// The implementation will also need a "Handle" method.
type EndpointHandler interface {
	EndpointPattern() string
}

// MiddlewareHandler allows for [[MiddlewareFunc]]s to be attached to an EndpointHandler.
type MiddlewareHandler interface {
	Middleware() []MiddlewareFunc
}

type EndpointWithTimeout interface {
	WithTimeout() time.Duration
}

// Validator should be used to validate JSON inputs for a request body.
type Validator interface {
	Validate() error
}
