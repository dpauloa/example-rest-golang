package usecase

import (
	"context"

	"dpauloa/example-rest-golang/domain"
)

//go:generate mockery -name CreatePhoneBook

type CreatePhoneBook interface {
	Execute(cxt context.Context, firstName, lastName, phoneNumber string) (*domain.PhoneBook, error)
}

type createPhoneBook struct {
	repo domain.PhoneBookRepo
}

func NewCreatePhoneBook(repo domain.PhoneBookRepo) CreatePhoneBook {
	return createPhoneBook{repo}
}

func (u createPhoneBook) Execute(cxt context.Context, firstName, lastName, phoneNumber string) (*domain.PhoneBook, error) {
	phoneBook, err := u.repo.CreatePhoneBook(cxt, firstName, lastName, phoneNumber)

	if err != nil {
		return nil, err
	}

	return phoneBook, nil
}

