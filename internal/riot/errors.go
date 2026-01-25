package riot

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	HTTPStatus        int
	Message           string
	RetryAfterSeconds *int
}

func (e *Error) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("riot error (%d): %s", e.HTTPStatus, e.Message)
	}
	return fmt.Sprintf("riot error (%d)", e.HTTPStatus)
}

type riotAPIError struct {
	Status struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"status"`
}

func parseRiotError(status int, body []byte, retryAfterSeconds *int) error {
	var re riotAPIError
	if err := json.Unmarshal(body, &re); err == nil && re.Status.Message != "" {
		return &Error{HTTPStatus: status, Message: re.Status.Message, RetryAfterSeconds: retryAfterSeconds}
	}
	// fallback if body isn't JSON
	msg := string(body)
	if msg == "" {
		msg = "upstream error"
	}
	return &Error{HTTPStatus: status, Message: msg, RetryAfterSeconds: retryAfterSeconds}
}
