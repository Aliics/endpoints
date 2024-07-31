package endpoints

type MiddlewareHandler interface {
	Middleware() []MiddlewareFunc
}

type EndpointHandler interface {
	EndpointPattern() string
}

type Validator interface {
	Validate() error
}
