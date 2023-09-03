package repository

import (
	"fmt"
	"github.com/wcodesoft/mosha-author-service/data"
)

// Repository represents the repository interface.
type Repository interface {
	AddAuthor(author data.Author) (string, error)
	ListAll() []data.Author
	UpdateAuthor(author data.Author) (data.Author, error)
	DeleteAuthor(id string) error
	GetAuthor(id string) (data.Author, error)
}

type repository struct {
	db               Database
	clientRepository ClientRepository
}

// AddAuthor adds a new author to the database.
func (s *repository) AddAuthor(author data.Author) (string, error) {
	return s.db.AddAuthor(author)
}

// ListAll returns all authors in the database.
func (s *repository) ListAll() []data.Author {
	return s.db.ListAll()
}

// UpdateAuthor updates an author in the database.
func (s *repository) UpdateAuthor(author data.Author) (data.Author, error) {
	return s.db.UpdateAuthor(author)
}

// DeleteAuthor deletes an author from the database.
func (s *repository) DeleteAuthor(id string) error {
	if err := s.deleteAuthorQuotes(id); err != nil {
		return err
	}
	return s.db.DeleteAuthor(id)
}

// GetAuthor returns an author from the database.
func (s *repository) GetAuthor(id string) (data.Author, error) {
	return s.db.GetAuthor(id)
}

// New creates a new repository.
func New(db Database, clientRepository ClientRepository) Repository {
	return &repository{
		db:               db,
		clientRepository: clientRepository,
	}
}

// deleteQuotes deletes all quotes from an author.
func (s *repository) deleteAuthorQuotes(id string) error {
	res, err := s.clientRepository.DeleteAuthorQuotes(id)
	if err != nil {
		return err
	}
	if !res {
		return fmt.Errorf("could not delete quotes from author with id: %s", id)
	}
	return nil
}
