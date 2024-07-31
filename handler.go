package endpoints

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
)

func bodyResponseHandler[Out any](output Out, err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			writeResponseErr(w, err)
			return
		}

		w.Header().Set(contentTypeHeader, "application/json")
		if err = json.NewEncoder(w).Encode(output); err != nil {
			// JSON writing failed. Reset the content-type and log.
			w.Header().Set(contentTypeHeader, "plain/text")
			slog.Error("could not write response", "err", err)
		}
	}
}

func errOnlyResponseHandler(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			writeResponseErr(w, err)
			return
		}
	}
}

func writeResponseErr(w http.ResponseWriter, err error) {
	handlerErr := &HandlerError{}
	if errors.As(err, handlerErr) {
		http.Error(w, handlerErr.Error(), handlerErr.StatusCode)
	} else {
		slog.Error("error when handling request", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
