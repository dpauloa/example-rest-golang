package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"dpauloa/example-rest-golang/domain/usecase"
)

func NewRouter(createPhoneBookUC usecase.CreatePhoneBook) http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)

	r.Get("/health", healthCheck)

	phoneResource := "/phones"
	r.Post(phoneResource, createPhoneBookHandler{createPhoneBookUC}.ServeHTTP)

	return r
}

func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
