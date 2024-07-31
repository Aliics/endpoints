package endpoints

import "net/http"

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request) func(next http.Handler)

type MiddlewareFromSlice struct {
	Middlewares []MiddlewareFunc
}

func (m MiddlewareFromSlice) Middleware() []MiddlewareFunc {
	return m.Middlewares
}

func executeMiddlewaresWithMiddlewares(w http.ResponseWriter, r *http.Request) func([]MiddlewareFunc, http.Handler) {
	return func(middlewares []MiddlewareFunc, handler http.Handler) {
		if len(middlewares) == 0 {
			// No middleware, just execute our handler.
			handler.ServeHTTP(w, r)
			return
		}

		var next http.Handler
		if len(middlewares) > 1 {
			next = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				executeMiddlewaresWithMiddlewares(w, r)(middlewares[1:], handler)
			})
		} else {
			next = handler
		}

		middlewares[0](w, r)(next)
	}
}
