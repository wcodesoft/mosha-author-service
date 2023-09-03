package repository

import (
	"github.com/wcodesoft/mosha-quote-service/data"
)

type FakeClientRepository struct {
	quotes   []data.Quote
	retValue bool
	retError error
	ClientRepository
}

func (f *FakeClientRepository) DeleteAuthorQuotes(_ string) (bool, error) {
	return f.retValue, f.retError
}

func (f *FakeClientRepository) SetDeleteAuthorQuotesReturn(value bool, err error) {
	f.retValue = value
	f.retError = err
}

func NewFakeClientRepository() *FakeClientRepository {
	return &FakeClientRepository{
		quotes:   []data.Quote{},
		retValue: true,
		retError: nil,
	}
}
