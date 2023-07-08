package service

import (
	"github.com/wcodesoft/mosha-author-service/data"
	"github.com/wcodesoft/mosha-author-service/repository"
)

// Service represents the service interface.
type Service interface {

	// CreateAuthor registers a new Author in the database.
	CreateAuthor(author data.Author) (string, error)

	// ListAll returns all authors in the database.
	ListAll() []data.Author

	// GetAuthor returns an author by id
	GetAuthor(id string) (data.Author, error)

	// DeleteAuthor deletes an author by id.
	DeleteAuthor(id string) error

	// UpdateAuthor updates an author.
	UpdateAuthor(author data.Author) (data.Author, error)
}

type service struct {
	repo repository.Repository
}

// New creates a new service.
func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

// CreateAuthor registers a new Author in the database.
func (s *service) CreateAuthor(author data.Author) (string, error) {
	return s.repo.AddAuthor(author)
}

// GetAuthor returns an author by id.
func (s *service) GetAuthor(id string) (data.Author, error) {
	return s.repo.GetAuthor(id)
}

// ListAll returns all authors in the database.
func (s *service) ListAll() []data.Author {
	return s.repo.ListAll()
}

// DeleteAuthor deletes an author by id.
func (s *service) DeleteAuthor(id string) error {
	return s.repo.DeleteAuthor(id)
}

// UpdateAuthor updates an author.
func (s *service) UpdateAuthor(author data.Author) (data.Author, error) {
	return s.repo.UpdateAuthor(author)
}
