package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"dpauloa/example-rest-golang/domain"
	"dpauloa/example-rest-golang/domain/usecase"

	"gopkg.in/validator.v2"
)

type createPhoneBookHandler struct {
	uc usecase.CreatePhoneBook
}

type phoneBookPayload struct {
	FirstName   string `json:"firstName" validate:"min=2,max=30"`
	LastName    string `json:"lastName" validate:"min=2,max=30"`
	PhoneNumber domain.PhoneNumber `json:"phoneNumber" validate:"min=15,max=15"`
}

type PhoneBookResponse struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

func (h createPhoneBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload phoneBookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		encodeJSONError(w, http.StatusBadRequest, errors.New("invalid body"))
		return
	}

	if errs := validator.Validate(payload); errs != nil {
		encodeJSONError(w, http.StatusBadRequest, errs)
		return
	}

	phoneNumberFormatted := payload.PhoneNumber.WithoutMask()

	response, err := h.uc.Execute(r.Context(), payload.FirstName, payload.LastName, phoneNumberFormatted)

	if err != nil {
		encodeJSONError(w, http.StatusBadRequest, err)
		return
	}

	encodeJSONStatus(w, http.StatusCreated, PhoneBookResponse{
		ID: response.ID,
		FirstName:   response.FirstName,
		LastName:    response.LastName,
		FullName:    response.FirstName + " " + response.LastName,
		PhoneNumber: response.PhoneNumber.WithMask(),
	})
}
