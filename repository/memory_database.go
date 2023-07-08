package repository

import (
	"fmt"
	"github.com/wcodesoft/mosha-author-service/data"
)

// inMemoryDatabase is a simple in-memory database.
type inMemoryDatabase struct {
	storage map[string]data.Author
}

// NewInMemoryDatabase creates a new InMemoryDatabase.
func NewInMemoryDatabase() Database {
	return &inMemoryDatabase{
		storage: make(map[string]data.Author),
	}
}

// AddAuthor adds a new author to the database.
func (db *inMemoryDatabase) AddAuthor(author data.Author) (string, error) {
	if _, ok := db.storage[author.ID]; ok {
		return "", fmt.Errorf("author %q already exist in database", author.ID)
	}
	db.storage[author.ID] = author
	return author.ID, nil
}

// ListAll returns all authors in the database.
func (db *inMemoryDatabase) ListAll() []data.Author {
	var authors []data.Author
	for _, v := range db.storage {
		authors = append(authors, v)
	}
	return authors
}

// UpdateAuthor updates an existing author in the database.
func (db *inMemoryDatabase) UpdateAuthor(author data.Author) (data.Author, error) {
	if _, ok := db.storage[author.ID]; !ok {
		return data.Author{}, fmt.Errorf("author %q do not exist in database", author.ID)
	}
	db.storage[author.ID] = author
	return db.storage[author.ID], nil
}

// DeleteAuthor deletes an existing author from the database.
func (db *inMemoryDatabase) DeleteAuthor(id string) error {
	if _, ok := db.storage[id]; !ok {
		return fmt.Errorf("author %q do not exist in database", id)
	}
	delete(db.storage, id)
	return nil
}

// GetAuthor returns an author from the database.
func (db *inMemoryDatabase) GetAuthor(id string) (data.Author, error) {
	if _, ok := db.storage[id]; !ok {
		return data.Author{}, fmt.Errorf("author %q do not exist in database", id)
	}
	return db.storage[id], nil
}
