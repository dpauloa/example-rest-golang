package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func jsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func encodeJSONError(w http.ResponseWriter, status int, err error) {
	jsonContentType(w)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: err.Error(),
	})
}

func encodeJSON(w http.ResponseWriter, v interface{}) {
	jsonContentType(w)
	json.NewEncoder(w).Encode(v)
}

func encodeJSONStatus(w http.ResponseWriter, status int, v interface{}) {
	jsonContentType(w)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func decodeJSON(body io.ReadCloser, v interface{}) {
	json.NewDecoder(body).Decode(v)
}
