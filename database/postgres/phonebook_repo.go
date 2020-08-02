package postgres

import (
	"context"
	"database/sql"
	"github.com/lib/pq"

	"dpauloa/example-rest-golang/domain"
)

type PhoneBookRepo struct {
	db *sql.DB
}

func NewPhoneBookRepo(db *sql.DB) PhoneBookRepo {
	return PhoneBookRepo{db}
}

func (r PhoneBookRepo) CreatePhoneBook(cxt context.Context, firstName, lastName, phoneNumber string) (*domain.PhoneBook, error) {
	query := "INSERT INTO phone_book (first_name, last_name, phone_number) VALUES ($1, $2, $3) RETURNING id"
	row:= r.db.QueryRowContext(cxt, query, firstName, lastName, phoneNumber)

	var id int64

	if err := row.Scan(&id); err != nil {
		if dbErr := err.(*pq.Error); dbErr.Constraint == "phone_book_phone_number_key" {
			return nil, domain.ErrPhoneNumberAlreadyExists
		}
		return nil, err
	}

	return &domain.PhoneBook{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		PhoneNumber: domain.PhoneNumber(phoneNumber),
	}, nil
}
