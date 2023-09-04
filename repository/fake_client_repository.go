package repository

import (
	"github.com/wcodesoft/mosha-quote-service/data"
)

type FakeClientRepository struct {
	quotes   []data.Quote
	retError error
	retRes   bool
	ClientRepository
}

func (f *FakeClientRepository) DeleteAuthorQuotes(_ string) (bool, error) {
	return f.retRes, f.retError
}

func (f *FakeClientRepository) SetDeleteAuthorQuotesReturn(ret bool, err error) {
	f.retError = err
	f.retRes = ret
}

func NewFakeClientRepository() *FakeClientRepository {
	return &FakeClientRepository{
		quotes:   []data.Quote{},
		retError: nil,
		retRes:   true,
	}
}
