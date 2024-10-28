package util

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HttpFunc func(w http.ResponseWriter, r *http.Request) error

func Make(f HttpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			slog.Error("An unexpected error occured", "Err:", err.Error())
			WriteJSON(w, http.StatusInternalServerError, "An unexpected error occured")
		}
	}
}

func WriteJSON(w http.ResponseWriter, code int, data any) error {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
