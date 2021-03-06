// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "dpauloa/example-rest-golang/domain"

	mock "github.com/stretchr/testify/mock"
)

// PhoneBookRepo is an autogenerated mock type for the PhoneBookRepo type
type PhoneBookRepo struct {
	mock.Mock
}

// CreatePhoneBook provides a mock function with given fields: cxt, firstName, lastName, phoneNumber
func (_m *PhoneBookRepo) CreatePhoneBook(cxt context.Context, firstName string, lastName string, phoneNumber string) (*domain.PhoneBook, error) {
	ret := _m.Called(cxt, firstName, lastName, phoneNumber)

	var r0 *domain.PhoneBook
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *domain.PhoneBook); ok {
		r0 = rf(cxt, firstName, lastName, phoneNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PhoneBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(cxt, firstName, lastName, phoneNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
