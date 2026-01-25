package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
)

type APIError struct {
	Code               string `json:"code"`
	Message            string `json:"message"`
	RetryAfterSeconds  *int   `json:"retryAfterSeconds,omitempty"`
}

type APIErrorResponse struct {
	Error APIError `json:"error"`
}

func writeAPIError(w http.ResponseWriter, status int, code, message string, retryAfterSeconds *int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(APIErrorResponse{
		Error: APIError{
			Code:              code,
			Message:           message,
			RetryAfterSeconds: retryAfterSeconds,
		},
	})
}

func writeMappedError(w http.ResponseWriter, err error) {
	var re *riot.Error
	if errors.As(err, &re) {
		switch re.HTTPStatus {
		case 401:
			writeAPIError(w, http.StatusUnauthorized, "RIOT_UNAUTHORIZED", "Riot API key invalid or missing", nil)
			return
		case 403:
			writeAPIError(w, http.StatusForbidden, "RIOT_FORBIDDEN", "Riot API key forbidden/expired", nil)
			return
		case 404:
			writeAPIError(w, http.StatusNotFound, "NOT_FOUND", "Riot account or resource not found", nil)
			return
		case 429:
			// pass Retry-After through to client
			writeAPIError(w, http.StatusTooManyRequests, "RATE_LIMITED", "Rate limited by Riot API. Retry later.", re.RetryAfterSeconds)
			return
		default:
			// Any other Riot status -> treat as bad gateway
			writeAPIError(w, http.StatusBadGateway, "UPSTREAM_ERROR", re.Message, re.RetryAfterSeconds)
			return
		}
	}

	// Non-Riot errors (network issues, json decode etc.)
	writeAPIError(w, http.StatusBadGateway, "UPSTREAM_ERROR", "Upstream request failed", nil)
}
