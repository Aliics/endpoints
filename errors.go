package endpoints

import (
	"net/http"
)

type HandlerError struct {
	StatusCode int
	Message    string
}

func NewHandlerError(code int, message string) HandlerError {
	return HandlerError{code, message}
}

func UnauthorizedError(message string) HandlerError {
	return NewHandlerError(http.StatusUnauthorized, message)
}

func PaymentRequiredError(message string) HandlerError {
	return NewHandlerError(http.StatusPaymentRequired, message)
}

func ForbiddenError(message string) HandlerError {
	return NewHandlerError(http.StatusForbidden, message)
}

func NotFoundError(message string) HandlerError {
	return NewHandlerError(http.StatusNotFound, message)
}

func MethodNotAllowedError(message string) HandlerError {
	return NewHandlerError(http.StatusMethodNotAllowed, message)
}

func NotAcceptableError(message string) HandlerError {
	return NewHandlerError(http.StatusNotAcceptable, message)
}

func ProxyAuthRequiredError(message string) HandlerError {
	return NewHandlerError(http.StatusProxyAuthRequired, message)
}

func RequestTimeoutError(message string) HandlerError {
	return NewHandlerError(http.StatusRequestTimeout, message)
}

func ConflictError(message string) HandlerError {
	return NewHandlerError(http.StatusConflict, message)
}

func GoneError(message string) HandlerError {
	return NewHandlerError(http.StatusGone, message)
}

func LengthRequiredError(message string) HandlerError {
	return NewHandlerError(http.StatusLengthRequired, message)
}

func PreconditionFailedError(message string) HandlerError {
	return NewHandlerError(http.StatusPreconditionFailed, message)
}

func RequestEntityTooLargeError(message string) HandlerError {
	return NewHandlerError(http.StatusRequestEntityTooLarge, message)
}

func RequestURITooLongError(message string) HandlerError {
	return NewHandlerError(http.StatusRequestURITooLong, message)
}

func UnsupportedMediaTypeError(message string) HandlerError {
	return NewHandlerError(http.StatusUnsupportedMediaType, message)
}

func RequestedRangeNotSatisfiableError(message string) HandlerError {
	return NewHandlerError(http.StatusRequestedRangeNotSatisfiable, message)
}

func ExpectationFailedError(message string) HandlerError {
	return NewHandlerError(http.StatusExpectationFailed, message)
}

func TeapotError(message string) HandlerError {
	return NewHandlerError(http.StatusTeapot, message)
}

func MisdirectedRequestError(message string) HandlerError {
	return NewHandlerError(http.StatusMisdirectedRequest, message)
}

func UnprocessableEntityError(message string) HandlerError {
	return NewHandlerError(http.StatusUnprocessableEntity, message)
}

func LockedError(message string) HandlerError {
	return NewHandlerError(http.StatusLocked, message)
}

func FailedDependencyError(message string) HandlerError {
	return NewHandlerError(http.StatusFailedDependency, message)
}

func TooEarlyError(message string) HandlerError {
	return NewHandlerError(http.StatusTooEarly, message)
}

func UpgradeRequiredError(message string) HandlerError {
	return NewHandlerError(http.StatusUpgradeRequired, message)
}

func PreconditionRequiredError(message string) HandlerError {
	return NewHandlerError(http.StatusPreconditionRequired, message)
}

func TooManyRequestsError(message string) HandlerError {
	return NewHandlerError(http.StatusTooManyRequests, message)
}

func RequestHeaderFieldsTooLargeError(message string) HandlerError {
	return NewHandlerError(http.StatusRequestHeaderFieldsTooLarge, message)
}

func UnavailableForLegalReasonsError(message string) HandlerError {
	return NewHandlerError(http.StatusUnavailableForLegalReasons, message)
}

func InternalServerErrorError(message string) HandlerError {
	return NewHandlerError(http.StatusInternalServerError, message)
}

func NotImplementedError(message string) HandlerError {
	return NewHandlerError(http.StatusNotImplemented, message)
}

func BadGatewayError(message string) HandlerError {
	return NewHandlerError(http.StatusBadGateway, message)
}

func ServiceUnavailableError(message string) HandlerError {
	return NewHandlerError(http.StatusServiceUnavailable, message)
}

func GatewayTimeoutError(message string) HandlerError {
	return NewHandlerError(http.StatusGatewayTimeout, message)
}

func HTTPVersionNotSupportedError(message string) HandlerError {
	return NewHandlerError(http.StatusHTTPVersionNotSupported, message)
}

func VariantAlsoNegotiatesError(message string) HandlerError {
	return NewHandlerError(http.StatusVariantAlsoNegotiates, message)
}

func InsufficientStorageError(message string) HandlerError {
	return NewHandlerError(http.StatusInsufficientStorage, message)
}

func LoopDetectedError(message string) HandlerError {
	return NewHandlerError(http.StatusLoopDetected, message)
}

func NotExtendedError(message string) HandlerError {
	return NewHandlerError(http.StatusNotExtended, message)
}

func NetworkAuthenticationRequiredError(message string) HandlerError {
	return NewHandlerError(http.StatusNetworkAuthenticationRequired, message)
}

func (h HandlerError) Error() string {
	return h.Message
}
