package domain

import "context"

type PhoneBookRepo interface {
	CreatePhoneBook(cxt context.Context, firstName, lastName, phoneNumber string) (*PhoneBook, error)
}