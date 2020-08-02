package domain

import (
	"errors"
	"strings"
)

type PhoneNumber string

type PhoneBook struct {
	ID          int64
	FirstName   string
	LastName    string
	PhoneNumber PhoneNumber
}

var (
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
)

func (pn PhoneNumber) WithoutMask() string {
	numberFormatted := strings.ReplaceAll(string(pn), "-", "")
	numberFormatted = strings.ReplaceAll(numberFormatted, "(", "")
	numberFormatted = strings.ReplaceAll(numberFormatted, ")", "")
	numberFormatted = strings.ReplaceAll(numberFormatted, " ", "")
	return numberFormatted
}

func (pn PhoneNumber) WithMask() string {
	return string("(" + pn[:2] + ") " + pn[2:7] + "-" + pn[7:])
}
