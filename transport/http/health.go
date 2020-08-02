package http

import (
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	encodeJSON(w, map[string]string{
		"status": "working",
	})
}
