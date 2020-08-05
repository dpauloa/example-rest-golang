package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"dpauloa/example-rest-golang/domain/usecase"
)

type Handler interface {
	http.Handler
}

func NewRouter(createPhoneBookUC usecase.CreatePhoneBook) http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)

	r.Get("/health", healthCheck)

	phoneResource := "/phones"
	r.Method(http.MethodPost, phoneResource, createPhoneBookHandler{createPhoneBookUC})

	return r
}

func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
