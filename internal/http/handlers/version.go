package handlers

import (
	"net/http"
)

func Version(v string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(v))
	}
}
