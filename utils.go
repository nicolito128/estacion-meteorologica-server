package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteString(w http.ResponseWriter, status int, s string, args ...any) (int, error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	return fmt.Fprintf(w, s, args...)
}

func WriteJSON(w http.ResponseWriter, status int, data any) (int, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	return WriteJSONBlob(w, status, b)
}

func WriteJSONBlob(w http.ResponseWriter, status int, data []byte) (int, error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return fmt.Fprintf(w, "%s", string(data))
}

func WriteError(w http.ResponseWriter, status int, err error) (int, error) {
	return WriteString(w, status, err.Error(), []any{}...)
}

func WriteJSONError(w http.ResponseWriter, status int, err error) (int, error) {
	return WriteJSONBlob(w, status, fmt.Appendf([]byte{}, `{"error": "%s"}`, err.Error()))
}
