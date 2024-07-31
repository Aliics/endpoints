package endpoints

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func responseHandler[Out any](output Out, err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			handlerErr := &HandlerError{}
			if errors.As(err, handlerErr) {
				http.Error(w, handlerErr.Error(), handlerErr.StatusCode)
			} else {
				slog.Error("error when handling request", "util", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(output); err != nil {
			// JSON writing failed. Reset the content-type and log.
			w.Header().Set("Content-Type", "plain/text")
			slog.Error("could not write response", "util", err)
		}
	}
}
