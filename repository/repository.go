package repository

import "authorservice/data"

// Repository represents the repository interface.
type Repository interface {
	AddAuthor(author data.Author) (string, error)
	ListAll() []data.Author
	UpdateAuthor(author data.Author) (data.Author, error)
	DeleteAuthor(id string) error
	GetAuthor(id string) (data.Author, error)
	AuthorExist(id string) bool
}

type repository struct {
	db Database
}

// AddAuthor adds a new author to the database.
func (s repository) AddAuthor(author data.Author) (string, error) {
	return s.db.AddAuthor(author)
}

// ListAll returns all authors in the database.
func (s repository) ListAll() []data.Author {
	return s.db.ListAll()
}

func (s repository) UpdateAuthor(author data.Author) (data.Author, error) {
	return s.db.UpdateAuthor(author)
}

func (s repository) DeleteAuthor(id string) error {
	return s.db.DeleteAuthor(id)
}

func (s repository) GetAuthor(id string) (data.Author, error) {
	return s.db.GetAuthor(id)
}

func (s repository) AuthorExist(id string) bool {
	return s.db.AuthorExist(id)
}

// New creates a new repository.
func New(db Database) Repository {
	return &repository{
		db: db,
	}
}
