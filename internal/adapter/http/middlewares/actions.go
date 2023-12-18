package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/account-api/infrastructure/logger"
)

func ResponseError(w http.ResponseWriter, r *http.Request, body []byte, err error) {
	var status int

	if errors.Is(err, errors.New("non existed resource")) || errors.Is(err, errors.New("validation failed")) {
		status = http.StatusBadRequest
	} else if errors.Is(err, errors.New("unauthorized")) {
		status = http.StatusUnauthorized
	} else if errors.Is(err, context.DeadlineExceeded) {
		status = http.StatusRequestTimeout
	} else {
		status = http.StatusInternalServerError
	}

	ResponseJSON(w, r, body, status, map[string]string{"error": err.Error()})
}

func ResponseJSON(w http.ResponseWriter, r *http.Request, body []byte, status int, payload interface{}) {
	for name, values := range r.Header {
		if name != "Content-Length" {
			for _, value := range values {
				w.Header().Set(name, value)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response, err := json.Marshal(payload)
	if err != nil {
		logger.Info(logger.ServerInfo, "Response payload could not be marshalled")
		response, _ = json.Marshal(map[string]string{"error": err.Error()})
	}

	w.Write(response)
	logger.Info(logger.ServerInfo, fmt.Sprintf("REQUEST %s:%s", r.Method, r.URL.String()))
	if body != nil {
		logger.Info(logger.ServerInfo, fmt.Sprintf("BODY %s", string(body)))
	}
	logger.Info(logger.ServerInfo, fmt.Sprintf("RESPONSE %d: %s", status, string(response)))
}
