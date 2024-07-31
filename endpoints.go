package endpoints

import (
	"context"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"net/url"
	"reflect"
)

const (
	endpointHandlerFuncName = "Handle"
)

// NewEndpointsMux will create a new mux with the given [[EndpointHandler]]s. This will also be setup with the
// appropriate routing, so they can serve HTTP requests.
func NewEndpointsMux(endpoints ...EndpointHandler) http.Handler {
	mux := http.NewServeMux()

	for _, endpoint := range endpoints {
		var middlewares []MiddlewareFunc
		if handler, ok := endpoint.(MiddlewareHandler); ok {
			middlewares = handler.Middleware()
		}

		method := mustFindEndpointHandleMethod(endpoint)

		mux.HandleFunc(endpoint.EndpointPattern(), func(w http.ResponseWriter, r *http.Request) {
			executeMiddlewaresWithMiddlewares(w, r)(middlewares, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				invokeEndpointHandleMethod(endpoint, method)(w, r)
			}))
		})
	}

	return mux
}

func invokeEndpointHandleMethod(endpoint EndpointHandler, method reflect.Method) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if endpointWithTimeout, ok := endpoint.(EndpointWithTimeout); ok {
			// We have a timeout configured on the endpoint.
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(r.Context(), endpointWithTimeout.WithTimeout())
			defer cancel()
		}

		arguments := []reflect.Value{
			reflect.ValueOf(endpoint),
		}
		for i := range method.Type.NumIn() {
			if i == 0 {
				// This will be a pointer to the "Endpoint".
				continue
			}

			fieldType := method.Type.In(i)

			var value any
			if fieldType.AssignableTo(reflect.TypeFor[context.Context]()) {
				value = ctx
			} else if fieldType.AssignableTo(reflect.TypeFor[url.Values]()) {
				value = r.URL.Query()
			} else if fieldType.AssignableTo(reflect.TypeFor[http.Header]()) {
				value = r.Header
			} else {
				// This is a type we would pull off the body. It will be JSON... because I said so.
				// We can convert the type into a value, and get the type we would like, but we will need to leverage
				// mapstructure or else we would have to make some bespoke solution to map a map[string]any to our
				// struct. I'm not doing that.
				// We also utilize the validator interface here.
				fieldValue := reflect.New(fieldType).Elem().Interface()
				var inputData any
				if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
					http.Error(w, "invalid json body", http.StatusBadRequest)
					return
				}

				if err := mapstructure.Decode(inputData, &fieldValue); err != nil {
					http.Error(w, "", http.StatusInternalServerError)
					return
				}

				if validated, ok := fieldValue.(Validator); ok {
					if err := validated.Validate(); err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
				}

				value = fieldValue
			}

			if value == nil {
				continue
			}

			arguments = append(arguments, reflect.ValueOf(value))
		}

		// Call our handler method in another goroutine to allow for ctx cancellations.
		outputChan := make(chan []reflect.Value)
		go func() { outputChan <- method.Func.Call(arguments) }()

		select {
		case output := <-outputChan:
			var err error
			if len(output) == 1 {
				if !output[0].IsZero() {
					err = output[0].Interface().(error)
				}

				errOnlyResponseHandler(err)(w, r)
			} else if len(output) == 2 {
				if !output[1].IsZero() {
					err = output[1].Interface().(error)
				}

				bodyResponseHandler(output[0].Interface(), err)(w, r)
			} else {
				panic("handle method must return a body or an error")
			}
		case <-ctx.Done():
			http.Error(w, "request timeout", http.StatusRequestTimeout)
		}
	}
}

func mustFindEndpointHandleMethod(endpoint EndpointHandler) reflect.Method {
	endpointType := reflect.TypeOf(endpoint)

	for i := range endpointType.NumMethod() {
		method := endpointType.Method(i)
		if method.Name == endpointHandlerFuncName {
			return method
		}
	}

	// e.g: func (e HelloEndpoint) Handle(in InputType) error { ... }
	panic(`endpoint must have a registered "Handle" method present`)
}
