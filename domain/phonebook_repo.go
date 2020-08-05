package domain

import "context"

//go:generate mockery -name PhoneBookRepo

type PhoneBookRepo interface {
	CreatePhoneBook(cxt context.Context, firstName, lastName, phoneNumber string) (*PhoneBook, error)
}