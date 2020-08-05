package usecase_test

import (
	"context"

	"dpauloa/example-rest-golang/domain"
	"dpauloa/example-rest-golang/domain/mocks"
	"dpauloa/example-rest-golang/domain/usecase"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Create PhoneBook", func() {
	var (
		ctx = context.Background()

		repo *mocks.PhoneBookRepo

		createPhoneBookUC usecase.CreatePhoneBook
	)

	const (
		id int64 = 1
		firstName = "danilo"
		lastName = "almeida"
		phoneNumber = "83900000000"
	)

	BeforeEach(func() {
		repo = &mocks.PhoneBookRepo{}
		createPhoneBookUC = usecase.NewCreatePhoneBook(repo)
	})

	Context("When values to create phone book is valid", func() {
		expectedPhoneBook := &domain.PhoneBook{
			ID:          id,
			FirstName:   firstName,
			LastName:    lastName,
			PhoneNumber: phoneNumber,
		}

		It("Should create a phone book", func() {
			repo.On("CreatePhoneBook", mock.Anything, firstName, lastName, phoneNumber).
				Return(expectedPhoneBook, nil)

			phoneBook, err := createPhoneBookUC.Execute(ctx, firstName, lastName, phoneNumber)

			Expect(err).To(BeNil())
			Expect(phoneBook.ID).To(Equal(id))
			Expect(phoneBook.FirstName).To(Equal(firstName))
			Expect(phoneBook.LastName).To(Equal(lastName))
			Expect(phoneBook.PhoneNumber).To(Equal(domain.PhoneNumber(phoneNumber)))
		})
	})
})
