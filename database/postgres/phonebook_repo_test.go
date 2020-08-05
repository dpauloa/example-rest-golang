package postgres_test

import (
	"context"
	"database/sql"

	"dpauloa/example-rest-golang/database/postgres"
	"dpauloa/example-rest-golang/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PhoneBookRepo", func() {
	var (
		ctx = context.Background()

		db *sql.DB
		mock sqlmock.Sqlmock

		repo postgres.PhoneBookRepo
	)

	const (
		firstName = "danilo"
		lastName = "almeida"
		phoneNumber = "83900000000"
	)

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		if err != nil {
			Fail("unable to create mock database")
		}

		repo = postgres.NewPhoneBookRepo(db)
	})

	AfterEach(func() {
		Expect(mock.ExpectationsWereMet()).To(BeNil(), "database expectations were not found")
		db.Close()
	})

	Describe("CreatePhoneBook", func() {
		Context("When there is no phone number with the given phone number", func() {
			It("Should create a phone book", func() {
				mock.ExpectQuery("INSERT INTO phone_book \\(first_name, last_name, phone_number\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
					WithArgs(firstName, lastName, phoneNumber).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

				phoneBook, err := repo.CreatePhoneBook(ctx, firstName, lastName, phoneNumber)

				Expect(err).To(BeNil())
				Expect(phoneBook.FirstName).To(Equal(firstName))
				Expect(phoneBook.LastName).To(Equal(lastName))
				Expect(phoneBook.PhoneNumber).To(Equal(domain.PhoneNumber(phoneNumber)))
			})
		})
		Context("When there is phone number with the given phone number", func() {
			It("Should not create a phone book", func() {
				mock.ExpectQuery("INSERT INTO phone_book \\(first_name, last_name, phone_number\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
					WithArgs(firstName, lastName, phoneNumber).
					WillReturnError(&pq.Error{Constraint: "phone_book_phone_number_key"})

				phoneBook, err := repo.CreatePhoneBook(ctx, firstName, lastName, phoneNumber)

				Expect(phoneBook).To(BeNil())
				Expect(err).To(Equal(domain.ErrPhoneNumberAlreadyExists))
			})
		})
	})
})
