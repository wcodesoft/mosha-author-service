package repository

import (
	"github.com/wcodesoft/mosha-quote-service/data"
)

type fakeClientRepository struct {
	quotes []data.Quote
}

func (f fakeClientRepository) DeleteAuthorQuotes(_ string) (bool, error) {
	return true, nil
}

func NewFakeClientRepository() ClientRepository {
	return &fakeClientRepository{
		quotes: []data.Quote{},
	}
}
