package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerId: 1001, Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateOfBirth: "2000-01-01", Status: true},
		{CustomerId: 1002, Name: "Rob", City: "New Delhi", Zipcode: "110011", DateOfBirth: "2000-01-01", Status: false},
	}

	return customerRepositoryMock{customers: customers}

}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerId == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
