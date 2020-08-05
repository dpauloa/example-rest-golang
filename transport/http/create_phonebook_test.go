package http_test

import (
	"bytes"
	"dpauloa/example-rest-golang/domain"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	gohttp "net/http"
	"net/http/httptest"

	"dpauloa/example-rest-golang/domain/usecase/mocks"
	"dpauloa/example-rest-golang/transport/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create PhoneBook", func() {
	var (
		createPhoneBookUC *mocks.CreatePhoneBook

		rw     *httptest.ResponseRecorder
		router http.Handler
	)

	const (
		id                  int64 = 1
		firstName                 = "danilo"
		lastName                  = "almeida"
		phoneNumber               = "83900000000"
		phoneNumberWithMask       = "(83) 90000-0000"
	)

	BeforeEach(func() {
		createPhoneBookUC = &mocks.CreatePhoneBook{}
		router = http.NewRouter(createPhoneBookUC)
	})

	Describe("#POST /phones", func() {
		Context("When request body is valid", func() {
			var (
				b, _ = json.Marshal(map[string]string{
					"firstName":   firstName,
					"lastName":    lastName,
					"phoneNumber": phoneNumberWithMask,
				})
			)

			Context("And phone has been successfully created", func() {
				BeforeEach(func() {
					expectedPhoneBook := &domain.PhoneBook{
						ID:          id,
						FirstName:   firstName,
						LastName:    lastName,
						PhoneNumber: phoneNumber,
					}

					createPhoneBookUC.On("Execute", mock.Anything, firstName, lastName, phoneNumber).
						Return(expectedPhoneBook, nil)

					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with phone book response and status code 201", func() {
					res := rw.Result()

					var body http.PhoneBookResponse
					json.NewDecoder(res.Body).Decode(&body)

					Expect(body.ID).To(Equal(id))
					Expect(body.FirstName).To(Equal(firstName))
					Expect(body.LastName).To(Equal(lastName))
					Expect(body.FullName).To(Equal(firstName + " " + lastName))
					Expect(body.PhoneNumber).To(Equal(domain.PhoneNumber(phoneNumber).WithMask()))
					Expect(res.StatusCode).To(Equal(201))
				})
			})

			Context("And phone already exists", func() {
				BeforeEach(func() {
					createPhoneBookUC.On("Execute", mock.Anything, firstName, lastName, phoneNumber).
						Return(nil, domain.ErrPhoneNumberAlreadyExists)

					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message error and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := domain.ErrPhoneNumberAlreadyExists.Error()
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})
		})

		Context("When request body is invalid", func() {
			Context("And body is empty", func() {
				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", nil)
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'invalid body' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "invalid body"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})
		})

		Context("When request body with attribute first name is invalid", func() {
			Context("And first name is empty", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   "",
						"lastName":    lastName,
						"phoneNumber": phoneNumberWithMask,
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'FirstName: less than min' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "FirstName: less than min"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})

			Context("And first name less than 2", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   "d",
						"lastName":    lastName,
						"phoneNumber": phoneNumberWithMask,
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'FirstName: less than min' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "FirstName: less than min"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})

			Context("And first name greater than 2", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   "danilo almeida danilo almeida danilo almeida danilo almeida danilo almeida",
						"lastName":    lastName,
						"phoneNumber": phoneNumberWithMask,
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'FirstName: greater than max' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "FirstName: greater than max"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})
		})

		Context("When request body with attribute last name is invalid", func() {
			Context("And last name is empty", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   firstName,
						"lastName":    "",
						"phoneNumber": phoneNumberWithMask,
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'LastName: less than min' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "LastName: less than min"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})

			Context("And last name less than 2", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   firstName,
						"lastName":    "d",
						"phoneNumber": phoneNumberWithMask,
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'LastName: less than min' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "LastName: less than min"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})

			Context("And last name greater than 2", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   firstName,
						"lastName":    "danilo almeida danilo almeida danilo almeida danilo almeida danilo almeida",
						"phoneNumber": phoneNumberWithMask,
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'LastName: greater than max' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "LastName: greater than max"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})
		})

		Context("When request body with attribute phone number is invalid", func() {
			Context("And phone number is empty", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   firstName,
						"lastName":    lastName,
						"phoneNumber": "",
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'PhoneNumber: less than min' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "PhoneNumber: less than min"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})

			Context("And phone number less than 15", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   firstName,
						"lastName":    firstName,
						"phoneNumber": "(83) 90000-000",
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'PhoneNumber: less than min' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "PhoneNumber: less than min"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})

			Context("And phone number greater than 15", func() {
				var (
					b, _ = json.Marshal(map[string]string{
						"firstName":   firstName,
						"lastName":    lastName,
						"phoneNumber": "(83) 90000-00000",
					})
				)

				BeforeEach(func() {
					req := httptest.NewRequest(gohttp.MethodPost, "/phones", bytes.NewBuffer(b))
					rw = httptest.NewRecorder()

					router.ServeHTTP(rw, req)
				})

				It("Should return body with message 'PhoneNumber: greater than max' and status code 400", func() {
					res := rw.Result()

					var body http.ErrorResponse
					json.NewDecoder(res.Body).Decode(&body)

					expectedMessage := "PhoneNumber: greater than max"
					Expect(body.Error).To(Equal(expectedMessage))
					Expect(res.StatusCode).To(Equal(400))
				})
			})
		})
	})
})
