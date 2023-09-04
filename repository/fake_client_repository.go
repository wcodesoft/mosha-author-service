package repository

import (
	"github.com/wcodesoft/mosha-quote-service/data"
)

type FakeClientRepository struct {
	quotes   []data.Quote
	retError error
	ClientRepository
}

func (f *FakeClientRepository) DeleteAuthorQuotes(_ string) error {
	return f.retError
}

func (f *FakeClientRepository) SetDeleteAuthorQuotesReturn(err error) {
	f.retError = err
}

func NewFakeClientRepository() *FakeClientRepository {
	return &FakeClientRepository{
		quotes:   []data.Quote{},
		retError: nil,
	}
}
