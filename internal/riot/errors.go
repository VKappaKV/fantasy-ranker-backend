package riot

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	HTTPStatus int
	Message    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("riot error (%d): %s", e.HTTPStatus, e.Message)
}

func parseRiotError(status int, body []byte) error {
	// Riot tipicamente manda { "status": { "message": "...", "status_code": ... } }
	var re RiotAPIError
	if err := json.Unmarshal(body, &re); err == nil && re.Status.Message != "" {
		return &Error{HTTPStatus: status, Message: re.Status.Message}
	}
	return &Error{HTTPStatus: status, Message: string(body)}
}
